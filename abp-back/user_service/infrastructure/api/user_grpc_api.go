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
	fmt.Print("user_service -> GetAll [api]\n")
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

func (handler *UserHandler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Println("In CreateUser grpc api")
	fmt.Print("Request.User: ")
	fmt.Println(request.User)
	user := mapNewUser(request.User)
	fmt.Print("user after mapping: ")
	fmt.Println(user)
	err := handler.service.Create(user)

	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		User: mapUser(user),
	}, nil
}

func (handler *UserHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Print("Request.Login: ")
	fmt.Println(request.Login)
	login := mapLogin(request.Login)
	fmt.Print("Login after mapping: ")
	fmt.Println(login)
	user, err := handler.service.Login(login.Username, login.Password)
	//var err error = nil
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		User: mapUser(user),
	}, nil
}

func (handler *UserHandler) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	fmt.Println("In UpdateUser grpc api")
	fmt.Print("Request.User: ")
	fmt.Println(request.User)
	user := mapUpdateUser(request.User)
	fmt.Print("user after mapping: ")
	fmt.Println(user)
	err := handler.service.UpdateUser(user)

	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{
		User: mapUser(user),
	}, nil
}
