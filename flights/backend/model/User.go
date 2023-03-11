package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Surname      string
	PhoneNumber  string
	UserRole     Role
	Email        string
	Password     string
	OwnedTickets []Ticket
}

func (user *User) BeforeCreate(scope *mongo.Database) error {
	user.ID = uuid.New()
	return nil
}
