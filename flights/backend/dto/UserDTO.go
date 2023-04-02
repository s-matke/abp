package dto

import "github.com/google/uuid"

type UserDTO struct {
	ID      uuid.UUID
	Name    string
	Surname string
	Email   string
	Role    string
}
