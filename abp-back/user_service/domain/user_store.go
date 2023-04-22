package domain

import "github.com/google/uuid"

type UserStore interface {
	Get(id uuid.UUID) (*Users, error)
	GetAll() (*[]Users, error)
	Insert(user *Users) error
	DeleteAll()
}
