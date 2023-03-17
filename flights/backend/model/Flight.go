package model

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	//ID             uuid.UUID `json:"id"`
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Departure      time.Time
	Origin         Location
	Destination    Location
	Price          float64
	AvailableSeats int
}

func (flight *Flight) BeforeCreate(*mongo.Database) error {
	flight.ID = primitive.NewObjectID()//uuid.New()
	return nil
}
