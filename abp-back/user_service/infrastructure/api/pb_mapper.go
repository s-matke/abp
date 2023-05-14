package api

import (
	"github.com/google/uuid"
	"github.com/s-matke/abp/abp-back/user_service/domain"

	pb "github.com/s-matke/abp/abp-back/common/proto/user_service"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:          user.Id.String(),
		Username:    user.Username,
		Name:        user.Name,
		Surname:     user.Surname,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
		Role:        pb.Role(user.Role),
		Address:     user.Address,
		City:        user.City,
		Country:     user.Country,
	}

	return userPb
}

func mapNewUser(userPb *pb.NewUser) *domain.User {
	user := &domain.User{

		//Id: uuid.New(),
		Username:    userPb.Username,
		Name:        userPb.Name,
		Surname:     userPb.Surname,
		PhoneNumber: userPb.PhoneNumber,
		Email:       userPb.Email,
		Password:    userPb.Password,
		Role:        domain.Role(userPb.Role),
		Address:     userPb.Address,
		City:        userPb.City,
		Country:     userPb.Country,
	}
	return user
}

func mapUpdateUser(userPb *pb.User) *domain.User {
	user := &domain.User{

		Id:          uuid.Must(uuid.Parse(userPb.Id)),
		Username:    userPb.Username,
		Name:        userPb.Name,
		Surname:     userPb.Surname,
		PhoneNumber: userPb.PhoneNumber,
		Email:       userPb.Email,
		Password:    userPb.Password,
		Role:        domain.Role(userPb.Role),
		Address:     userPb.Address,
		City:        userPb.City,
		Country:     userPb.Country,
	}
	return user
}
