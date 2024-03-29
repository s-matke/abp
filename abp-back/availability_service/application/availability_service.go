package application

import (
	"github.com/s-matke/abp/abp-back/availability_service/domain"
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

func (service *AvailabilityService) GetByAccommodation(id string) ([]*domain.Availability, error) {
	return service.store.GetByAccommodation(id)
}

func (service *AvailabilityService) CreateAvailability(availability *domain.Availability) error {
	return service.store.Insert(availability)
}

func (service *AvailabilityService) DeleteByData(availability *domain.Availability) error {
	availabilities, err := service.store.GetByAccommodation(availability.AccommodationId.Hex())

	if err != nil {
		return err
	}

	for _, oldAvailability := range availabilities {
		if availability.StartDate.Compare(oldAvailability.StartDate) == 0 && availability.EndDate.Compare(oldAvailability.EndDate) == 0 {
			return service.store.DeleteByData(oldAvailability)
		}
	}

	return nil
}

func (service *AvailabilityService) GetAllUnavailable(availability *domain.Availability) ([]*domain.Availability, error) {
	availabilities, err := service.store.GetAll()

	if err != nil {
		return nil, err
	}

	var unavailable []*domain.Availability

	for _, existing := range availabilities {
		if availability.StartDate.Before(existing.EndDate) && existing.StartDate.Before(availability.EndDate) {
			unavailable = append(unavailable, existing)
		}
	}

	return unavailable, err
}
