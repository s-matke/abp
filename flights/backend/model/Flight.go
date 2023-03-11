package model

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Flight struct {
	ID             uuid.UUID `json:"id"`
	Departure      time.Time
	Origin         Location
	Destination    Location
	Price          float64
	AvailableSeats int
}

func (flight *Flight) BeforeCreate(*mongo.Database) error {
	flight.ID = uuid.New()
	return nil
}
