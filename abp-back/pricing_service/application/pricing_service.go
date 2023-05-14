package application

import (
	"github.com/s-matke/abp/abp-back/pricing_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PricingService struct {
	store domain.PricingStore
}

func NewPricingService(store domain.PricingStore) *PricingService {
	return &PricingService{
		store: store,
	}
}

func (service *PricingService) Get(id primitive.ObjectID) (*domain.Pricing, error) {
	return service.store.Get(id)
}

func (service *PricingService) GetAll() ([]*domain.Pricing, error) {
	return service.store.GetAll()
}

func (service *PricingService) GetByAccommodation(id string) (*domain.Pricing, error) {
	return service.store.GetByAccommodation(id)
}
