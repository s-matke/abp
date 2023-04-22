package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id          uuid.UUID `json:"id" gorm:"primaryKey;not null;unique"`
	Username    string    `json:"username" gorm:"unique;not null"`
	Name        string    `json:"name" gorm:"not null"`
	Surname     string    `json:"surname" gorm:"not null"`
	PhoneNumber string    `json:"phonenumber"`
	Email       string    `json:"email" gorm:"not null"`
	Password    string    `json:"password" gorm:"not null"`
	Role        Role      `json:"role" gorm:"not null"`
	Location    Location  `json:"location"`
}

type Role int

const (
	GUEST Role = iota
	HOST
)

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.Id = uuid.New()
	return nil
}
