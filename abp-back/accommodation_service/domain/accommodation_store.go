package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccommodationStore interface {
	Get(id primitive.ObjectID) (*Accommodation, error)
	GetAll() ([]*Accommodation, error)
	Insert(accommodation *Accommodation) error
	DeleteAll()
	GetByHost(id string) ([]*Accommodation, error)
	SearchAccommodation(availableSeats int32, destination string, ids []primitive.ObjectID) ([]*Accommodation, error)
}
