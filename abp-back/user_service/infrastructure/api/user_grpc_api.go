package api

import (
	"context"
	"fmt"

	"github.com/s-matke/abp/abp-back/user_service/domain"

	"github.com/s-matke/abp/abp-back/user_service/application"

	pb "github.com/s-matke/abp/abp-back/common/proto/user_service"

	"github.com/google/uuid"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	print(id)
	fmt.Printf("type id: %T\n", id)

	uid, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	var user *domain.User
	user, err = handler.service.Get(uid)

	if err != nil {
		return nil, err
	}

	userPb := mapUser(user)

	response := &pb.GetResponse{
		User: userPb,
	}

	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	users, err := handler.service.GetAll()

	if err != nil || *users == nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}

	for _, user := range *users {
		current := mapUser(&user)
		response.Users = append(response.Users, current)
	}

	return response, nil
}
