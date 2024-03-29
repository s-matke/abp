package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	accommodation "github.com/s-matke/abp/abp-back/common/proto/accommodation_service"
	availability "github.com/s-matke/abp/abp-back/common/proto/availability_service"
	pricing "github.com/s-matke/abp/abp-back/common/proto/pricing_service"
	"github.com/s-matke/abp/abp-back/reservation_service/domain"
	"github.com/s-matke/abp/abp-back/reservation_service/infrastructure/persistence"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ReservationService struct {
	store                    domain.ReservationStore
	accommodationServiceAddr string
	pricingServiceAddr       string
	availabilityAddr         string
}

func NewReservationService(store domain.ReservationStore, accommodationServiceAddr, pricingServiceAddr, availabilityAddr string) *ReservationService {
	return &ReservationService{
		store:                    store,
		accommodationServiceAddr: accommodationServiceAddr,
		pricingServiceAddr:       pricingServiceAddr,
		availabilityAddr:         availabilityAddr,
	}
}

func (service *ReservationService) GetAll() ([]*domain.Reservation, error) {
	return service.store.GetAll()
}

func (service *ReservationService) GetByAccommodation(id primitive.ObjectID) ([]*domain.Reservation, error) {
	return service.store.GetByAccommodation(id)
}

func (service *ReservationService) GetByGuest(id string) ([]*domain.Reservation, error) {
	return service.store.GetByGuest(id)
}

func (service *ReservationService) GetAllPendingByAccommodation(id primitive.ObjectID) ([]*domain.Reservation, error) {
	return service.store.GetAllPendingByAccommodation(id)
}

func (service *ReservationService) Insert(reservation *domain.Reservation) error {
	return service.store.Insert(reservation)
}

func (service *ReservationService) GetCancelledAmount(id string) int32 {
	return service.store.GetCancelledAmount(id)
}

func (service *ReservationService) CancelReservation(id primitive.ObjectID) error {

	reservation, err := service.store.Get(id)

	if err != nil {
		return err
	}

	err = service.store.CancelReservation(id)

	if err != nil {
		return nil
	}

	if reservation.Status == domain.PENDING {
		return nil
	}

	availabilityClient := persistence.NewAvailabilityService(service.availabilityAddr)

	availabilityResponse, err := availabilityClient.DeleteByData(context.TODO(), &availability.DeleteByDataRequest{
		AccommodationId: reservation.AccommodationId.Hex(),
		StartDate:       timestamppb.New(reservation.StartDate),
		EndDate:         timestamppb.New(reservation.EndDate),
	})

	print(availabilityResponse)

	if err != nil {
		return err
	}

	return nil

}

func (service *ReservationService) ConfirmReservation(id primitive.ObjectID) ([]*domain.Reservation, error) {
	reservation, err := service.store.ConfirmReservation(id)

	if err != nil {
		return nil, err
	}

	fmt.Println("Izmeni na BOOKED")

	availabilityClient := persistence.NewAvailabilityService(service.availabilityAddr)

	accommodation_id := reservation.AccommodationId
	startDate := reservation.StartDate
	endDate := reservation.EndDate

	availabilityResponse, err := availabilityClient.CreateAvailability(context.TODO(), &availability.CreateAvailabilityRequest{
		Availability: &availability.NewAvailability{
			AccommodationId: accommodation_id.Hex(),
			StartDate:       timestamppb.New(startDate),
			EndDate:         timestamppb.New(endDate),
		},
	})

	if err != nil {
		return nil, err
	}

	print(availabilityResponse)

	fmt.Println("Prosao create availability")

	pendingReservations, err := service.GetAllPendingByAccommodation(accommodation_id)

	if err != nil {
		return nil, err
	}

	var reservationsToSave []*domain.Reservation
	var reservationsToDelete []primitive.ObjectID

	for _, reservation := range pendingReservations {
		if startDate.Before(reservation.EndDate) && reservation.StartDate.Before(endDate) {
			reservationsToDelete = append(reservationsToDelete, reservation.Id)
		} else {
			reservationsToSave = append(reservationsToSave, reservation)
		}
	}

	fmt.Println("Broj za brisanje: ", len(reservationsToDelete))

	if len(reservationsToDelete) != 0 {
		err = service.store.DeleteByIds(reservationsToDelete)
		if err != nil {
			return nil, err
		}
	}

	return reservationsToSave, nil
}

func (service *ReservationService) CreateReservation(reservation *domain.Reservation) (*domain.Reservation, error) {
	accommodationClient := persistence.NewAccommodationService(service.accommodationServiceAddr)

	accommodationResponse, err := accommodationClient.Get(context.TODO(), &accommodation.GetRequest{Id: reservation.AccommodationId.Hex()})

	if err != nil {
		return nil, err
	}

	if accommodationResponse.Accommodation.AutomaticReservation {
		reservation.Status = domain.BOOKED
	} else {
		reservation.Status = domain.PENDING
	}

	availabilityClient := persistence.NewAvailabilityService(service.availabilityAddr)

	availabilityResponse, err := availabilityClient.GetByAccommodation(context.TODO(), &availability.GetByAccommodationRequest{Id: accommodationResponse.Accommodation.Id})

	if err != nil {
		return nil, err
	}

	for _, availability := range availabilityResponse.Availabilities {
		existingStartDate := time.Date(
			availability.StartDate.AsTime().Year(),
			availability.StartDate.AsTime().Month(),
			availability.StartDate.AsTime().Day(),
			0, 0, 0, 0, availability.StartDate.AsTime().Location(),
		)
		existingEndDate := time.Date(
			availability.EndDate.AsTime().Year(),
			availability.EndDate.AsTime().Month(),
			availability.EndDate.AsTime().Day(),
			0, 0, 0, 0, availability.EndDate.AsTime().Location(),
		)

		if reservation.StartDate.Before(existingEndDate) && existingStartDate.Before(reservation.EndDate) {
			return nil, errors.New("date overlapping with existing reservation")
		}
	}

	pricingClient := persistence.NewPricingClient(service.pricingServiceAddr)

	pricingResponse, err := pricingClient.CalculatePrice(context.TODO(), &pricing.CalculatePriceRequest{
		Id:        accommodationResponse.Accommodation.Id,
		StartDate: timestamppb.New(reservation.StartDate),
		EndDate:   timestamppb.New(reservation.EndDate),
		NumPeople: int32(reservation.NumOfGuests),
	})

	if err != nil {
		return nil, err
	}

	reservation.Price = pricingResponse.Price
	reservation.Id = primitive.NewObjectID()

	err = service.Insert(reservation)

	if err != nil {
		return nil, err
	}

	availabilityNew := availability.NewAvailability{
		AccommodationId: reservation.AccommodationId.Hex(),
		StartDate:       timestamppb.New(reservation.StartDate),
		EndDate:         timestamppb.New(reservation.EndDate),
	}

	if reservation.Status == domain.PENDING {
		return reservation, nil
	} else {

		_, err = availabilityClient.CreateAvailability(context.TODO(), &availability.CreateAvailabilityRequest{
			Availability: &availabilityNew,
		})

		if err != nil {
			return nil, err
		}

		return reservation, nil
	}
}
