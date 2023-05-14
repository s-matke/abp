package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PricingStore interface {
	GetByAccommodation(id string) (*Pricing, error)
	Get(id primitive.ObjectID) (*Pricing, error)
	GetAll() ([]*Pricing, error)
	Insert(pricing *Pricing) error
	DeleteAll()
}
