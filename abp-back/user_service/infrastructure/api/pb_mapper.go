package api

import (
	pb "common/proto/user_service"
	"user_service/domain"
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
		Location: &pb.Location{
			City:    user.Location.City,
			Country: user.Location.Country,
		},
	}

	return userPb
}
