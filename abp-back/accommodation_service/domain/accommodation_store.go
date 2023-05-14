package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccommodationStore interface {
	Get(id primitive.ObjectID) (*Accommodation, error)
	GetAll() ([]*Accommodation, error)
	Insert(accommodation *Accommodation) error
	DeleteAll()
}
