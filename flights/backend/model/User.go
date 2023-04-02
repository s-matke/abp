package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `bson:"id,omitempty" unique:"true"`
	Name         string
	Surname      string
	PhoneNumber  string `bson:"phonenumber,omitempty" unique:"true"`
	UserRole     Role
	Email        string `bson:"email" unique:"true"`
	Password     string
	OwnedTickets []Ticket
	Location     Location
}
