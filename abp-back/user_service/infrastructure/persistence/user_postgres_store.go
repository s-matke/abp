package persistence

import (
	"github.com/s-matke/abp/abp-back/user_service/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPostgresStore struct {
	db *gorm.DB
}

func NewUserPostgresStore(db *gorm.DB) (domain.UserStore, error) {
	err := db.AutoMigrate((&domain.Users{}))
	if err != nil {
		return nil, err
	}
	return &UserPostgresStore{
		db: db,
	}, nil
}

func (store *UserPostgresStore) Get(id uuid.UUID) (*domain.Users, error) {
	var user domain.Users
	// result := store.db.First(&user, "id = ?", id)
	result := store.db.Where("id = ?", id).First(&user) // isto ko ovo gore
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (store *UserPostgresStore) GetAll() (*[]domain.Users, error) {
	var users []domain.Users
	result := store.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (store *UserPostgresStore) Insert(user *domain.Users) error {
	result := store.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *UserPostgresStore) DeleteAll() {
	store.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.Users{})
}
