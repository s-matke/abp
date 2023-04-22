package api

import (
	"github.com/s-matke/abp/abp-back/user_service/domain"

	pb "github.com/s-matke/abp/abp-back/common/proto/user_service"
)

func mapUser(user *domain.Users) *pb.Users {
	userPb := &pb.Users{
		Id:          user.Id.String(),
		Username:    user.Username,
		Name:        user.Name,
		Surname:     user.Surname,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
		// Role:        pb.Role(user.Role),
		// Location: &pb.Location{
		// 	City:    user.Location.City,
		// 	Country: user.Location.Country,
		// },
	}

	return userPb
}
