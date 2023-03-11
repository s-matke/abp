package model

import "github.com/google/uuid"

type User struct {
	Id          uuid.UUID
	Name        string
	Surname     string
	PhoneNumber string
	UserRole    Role
	Email       string
	Password    string
}
