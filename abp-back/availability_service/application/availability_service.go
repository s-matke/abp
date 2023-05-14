package application

import (
	"github.com/s-matke/abp/abp-back/availability_service/domain"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type AvailabilityService struct {
	store domain.AvailabilityStore
}

func NewAvailabilityService(store domain.AvailabilityStore) *AvailabilityService {
	return &AvailabilityService{
		store: store,
	}
}

func (service *AvailabilityService) GetAll() ([]*domain.Availability, error) {
	return service.store.GetAll()
}

func (service *AvailabilityService) CreateAvailability(availability *domain.Availability) error {
	return service.store.Insert(availability)
}

func (service *AvailabilityService) GetAllUnavailable(startDate *timestamp.Timestamp, endDate  *timestamp.Timestamp ) ([]*string, error) {
	return service.store.GetAllUnavailable(startDate, endDate)
}
