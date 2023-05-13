package application

import (
	"fmt"

	"github.com/s-matke/abp/abp-back/user_service/domain"

	"github.com/google/uuid"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(id uuid.UUID) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetAll() (*[]domain.User, error) {
	fmt.Print("user_service -> GetAll [application]\n")
	return service.store.GetAll()
}
