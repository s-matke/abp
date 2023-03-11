package service

import "flight/repository"

type UserService struct {
	UserRepository *repository.UserRepository
}
