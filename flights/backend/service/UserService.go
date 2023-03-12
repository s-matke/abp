package service

import (
	"errors"
	"flight/model"
	"flight/repository"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (service *UserService) Create(user *model.User) error {

	userExists, _ := service.UserRepository.FindByEmail(user.Email)

	if userExists != nil {
		return errors.New("user with given email already exists")
	}

	user.ID = uuid.New()
	user.UserRole = model.REGULAR

	err := service.UserRepository.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}
