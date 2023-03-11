package model

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Ticket struct {
	ID       uuid.UUID
	Flight   Flight
	IssuedAt time.Time
}

func (ticket *Ticket) BeforeCreate(*mongo.Database) error {
	ticket.ID = uuid.New()
	return nil
}
