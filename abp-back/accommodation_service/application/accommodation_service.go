package application

import (
	"github.com/s-matke/abp/abp-back/accommodation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	store domain.AccommodationStore
}

func NewAccommodationService(store domain.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}

func (service *AccommodationService) Get(id primitive.ObjectID) (*domain.Accommodation, error) {
	return service.store.Get(id)
}

func (service *AccommodationService) GetAll() ([]*domain.Accommodation, error) {
	return service.store.GetAll()
}

func (service *AccommodationService) CreateAccommodation(accommodation *domain.Accommodation) error {
	return service.store.Insert(accommodation)
}

func (service *AccommodationService) GetByHost(id string) ([]*domain.Accommodation, error) {
	return service.store.GetByHost(id)
}
