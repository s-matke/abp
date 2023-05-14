package application

import (
	"github.com/s-matke/abp/abp-back/reservation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationService struct {
	store domain.ReservationStore
}

func NewReservationService(store domain.ReservationStore) *ReservationService {
	return &ReservationService{
		store: store,
	}
}

func (service *ReservationService) GetAll() ([]*domain.Reservation, error) {
	return service.store.GetAll()
}

func (service *ReservationService) GetByAccommodation(id primitive.ObjectID) ([]*domain.Reservation, error) {
	return service.store.GetByAccommodation(id)
}

func (service *ReservationService) Insert(reservation *domain.Reservation) error {
	return service.store.Insert(reservation)
}
