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

func (service *UserService) Create(user *domain.User) error {
	/*checkUser, err := service.store.GetByUsername(user.Username)
	if err != nil {
	    if errors.Is(err, gorm.ErrRecordNotFound) {
			// handle record not found error
		} else {
			return err
		}
	}
	fmt.Print("Get user by username: ")
	fmt.Println(checkUser)
	if checkUser != nil  {
		return fmt.Errorf("user with username %s already exists", user.Username)
	}

	checkUser, err = service.store.GetByEmail(user.Email)
	if err != nil {
	    return err
	}
	if checkUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	fmt.Print("Get user by email: ")
	fmt.Println(checkUser)*/
	//TODO: Optional
	err := service.store.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateUser(user *domain.User) error {
	err := service.store.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
