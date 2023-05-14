package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status int

const (
	BOOKED Status = iota
	CANCELLED
	PENDING
)

type Reservation struct {
	Id              primitive.ObjectID `bson:"_id"`
	AccommodationId primitive.ObjectID `bson:"accommodation_id"`
	GuestId         string             `bson:"guest_id"`
	StartDate       time.Time          `bson:"start_date"`
	EndDate         time.Time          `bson:"end_date"`
	NumOfGuests     int                `bson:"num_of_guests"`
	Price           float32            `bson:"price"`
	Status          Status             `bson:"status"`
}
