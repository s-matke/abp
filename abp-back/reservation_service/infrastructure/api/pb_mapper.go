package api

import (
	"time"

	pb "github.com/s-matke/abp/abp-back/common/proto/reservation_service"
	"github.com/s-matke/abp/abp-back/reservation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapReservation(reservation *domain.Reservation) *pb.Reservation {
	reservationPb := &pb.Reservation{
		Id:              reservation.Id.Hex(),
		AccommodationId: reservation.AccommodationId.Hex(),
		GuestId:         reservation.GuestId,
		StartDate:       timestamppb.New(reservation.StartDate),
		EndDate:         timestamppb.New(reservation.EndDate),
		NumOfGuests:     int32(reservation.NumOfGuests),
		Price:           reservation.Price,
		Status:          mapStatus(reservation.Status),
	}

	return reservationPb
}

func mapStatus(status domain.Status) pb.Reservation_Status {
	switch status {
	case domain.BOOKED:
		return pb.Reservation_BOOKED
	case domain.PENDING:
		return pb.Reservation_PENDING
	}
	return pb.Reservation_CANCELLED
}

func mapNewReservation(request *pb.CreateReservationRequest) *domain.Reservation {

	accommodation_id, err := primitive.ObjectIDFromHex(request.Reservation.AccommodationId)

	if err != nil {
		return nil
	}

	reservation := &domain.Reservation{
		AccommodationId: accommodation_id,
		GuestId:         request.Reservation.GuestId,
		StartDate: time.Date(
			request.Reservation.StartDate.AsTime().Year(),
			request.Reservation.StartDate.AsTime().Month(),
			request.Reservation.StartDate.AsTime().Day(),
			0, 0, 0, 0,
			request.Reservation.StartDate.AsTime().Location(),
		),
		EndDate: time.Date(
			request.Reservation.EndDate.AsTime().Year(),
			request.Reservation.EndDate.AsTime().Month(),
			request.Reservation.EndDate.AsTime().Day(),
			0, 0, 0, 0,
			request.Reservation.EndDate.AsTime().Location(),
		),
		NumOfGuests: int(request.Reservation.NumOfGuests),
	}
	return reservation
}
