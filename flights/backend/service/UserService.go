package service

import (
	"errors"
	"flight/model"
	"flight/repository"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (service *UserService) Create(user *model.User) error {

	userExists, _ := service.FindByEmail(user.Email)

	if userExists != nil {
		return errors.New("user with given email already exists")
	}

	user.ID = uuid.New()
	user.UserRole = model.REGULAR
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(bytes)

	err := service.UserRepository.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (service *UserService) FindByEmail(email string) (*model.User, error) {
	var user model.User
	userExists, err := service.UserRepository.FindByEmail(email)

	if err != nil {
		return nil, errors.New("user not found")
	}

	bsonBytes, _ := bson.Marshal(userExists)
	_ = bson.Unmarshal(bsonBytes, &user)

	return &user, nil
}
