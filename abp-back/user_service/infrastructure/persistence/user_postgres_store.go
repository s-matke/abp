package persistence

import (
	"errors"
	"fmt"

	//"context"
	"github.com/s-matke/abp/abp-back/user_service/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPostgresStore struct {
	db *gorm.DB
}

func NewUserPostgresStore(db *gorm.DB) (domain.UserStore, error) {
	err := db.AutoMigrate((&domain.User{}))
	if err != nil {
		return nil, err
	}
	return &UserPostgresStore{
		db: db,
	}, nil
}

func (store *UserPostgresStore) Get(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	// result := store.db.First(&user, "id = ?", id)
	result := store.db.Where("id = ?", id).First(&user) // isto ko ovo gore
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (store *UserPostgresStore) GetAll() (*[]domain.User, error) {
	fmt.Print("user_service -> GetAll [persistance]\n")
	var users []domain.User
	result := store.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (store *UserPostgresStore) Insert(user *domain.User) error {
	result := store.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *UserPostgresStore) DeleteAll() {
	store.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.User{})
}

func (store *UserPostgresStore) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	result := store.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (store *UserPostgresStore) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := store.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (store *UserPostgresStore) UpdateUser(user *domain.User) error {
	var newUser domain.User
	tx := store.db.Where("id = ?", user.Id).Model(&newUser).Updates(domain.User{

		Name:        user.Name,
		Surname:     user.Surname,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
		Address:     user.Address,
		City:        user.City,
		Country:     user.Country,
	})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected != 1 {
		return errors.New("update error")
	}
	return nil
}
