package model

import (
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	Id             uuid.UUID `json:"id"`
	Departure      time.Time
	Origin         Location
	Destination    Location
	Price          float64
	availableSeats int
}
