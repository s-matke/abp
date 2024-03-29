package domain

import "github.com/google/uuid"

type UserStore interface {
	Get(id uuid.UUID) (*User, error)
	GetAll() (*[]User, error)
	Insert(user *User) error
	DeleteAll()
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	Login(username string, email string) (*User, error)
}
