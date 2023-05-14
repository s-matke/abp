package application

import (
	"context"
	"fmt"

	"github.com/s-matke/abp/abp-back/accommodation_service/domain"
	"github.com/s-matke/abp/abp-back/accommodation_service/infrastructure/persistence"
	availability "github.com/s-matke/abp/abp-back/common/proto/availability_service"
	pricing "github.com/s-matke/abp/abp-back/common/proto/pricing_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	store                   domain.AccommodationStore
	pricingServiceAddr      string
	availabilityServiceAddr string
}

func NewAccommodationService(store domain.AccommodationStore, pricingServiceAddr string, availabilityServiceAddr string) *AccommodationService {
	return &AccommodationService{
		store:                   store,
		pricingServiceAddr:      pricingServiceAddr,
		availabilityServiceAddr: availabilityServiceAddr,
	}
}

func (service *AccommodationService) Get(id primitive.ObjectID) (*domain.Accommodation, error) {
	return service.store.Get(id)
}

func (service *AccommodationService) GetAll() ([]*domain.Accommodation, error) {
	pricingClient := persistence.NewPricingClient(service.pricingServiceAddr)
	availabilityClient := persistence.NewAvailabilityClient(service.availabilityServiceAddr)
	pricingResponse, err := pricingClient.GetAll(context.TODO(), &pricing.GetAllRequest{})

	if err != nil {
		return nil, err
	}
	availabilityResponse, err := availabilityClient.GetAll(context.TODO(), &availability.GetAllRequest{})

	if err != nil {
		return nil, err
	}

	print("=======================")
	// ovde ispise adresu neku
	print(pricingResponse)
	// Unutar ovog for loop-a ispisuje svaki pricing kako treba (slika na fejsu)
	for _, pricing := range pricingResponse.Pricings {
		fmt.Println("Jedan pricing: ")
		fmt.Println(pricing)
	}

	for _, availability := range availabilityResponse.Availabilities {
		fmt.Println(availability)
	}

	print("=======================")

	return service.store.GetAll()
}

func (service *AccommodationService) CreateAccommodation(accommodation *domain.Accommodation) error {
	return service.store.Insert(accommodation)
}

func (service *AccommodationService) GetByHost(id string) ([]*domain.Accommodation, error) {
	return service.store.GetByHost(id)
}
