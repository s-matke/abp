package domain

import "github.com/google/uuid"

type UserStore interface {
	Get(id uuid.UUID) (*User, error)
	GetAll() (*[]User, error)
	Insert(user *User) error
	DeleteAll()
}
