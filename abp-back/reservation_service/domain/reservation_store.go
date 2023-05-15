package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReservationStore interface {
	GetAll() ([]*Reservation, error)
	GetByAccommodation(id primitive.ObjectID) ([]*Reservation, error)
	Insert(reservation *Reservation) error
	DeleteAll()
	GetAllPendingByAccommodation(id primitive.ObjectID) ([]*Reservation, error)
	GetCancelledAmount(id string) int32
	ConfirmReservation(id primitive.ObjectID) (*Reservation, error)
	DeleteByIds(id []primitive.ObjectID) error
}
