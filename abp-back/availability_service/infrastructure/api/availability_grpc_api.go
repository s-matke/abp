package api

import (
	"context"

	"github.com/s-matke/abp/abp-back/availability_service/application"
	pb "github.com/s-matke/abp/abp-back/common/proto/availability_service"
)

type AvailabilityHandler struct {
	pb.UnimplementedAvailabilityServiceServer
	service *application.AvailabilityService
}

func NewAvailabilityHandler(service *application.AvailabilityService) *AvailabilityHandler {
	return &AvailabilityHandler{
		service: service,
	}
}

func (handler *AvailabilityHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	availabilities, err := handler.service.GetAll()

	if err != nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		Availabilities: []*pb.Availability{},
	}

	for _, availability := range availabilities {
		current := mapAvailability(availability)
		response.Availabilities = append(response.Availabilities, current)
	}

	return response, nil
}

func (handler *AvailabilityHandler) CreateAvailability(ctx context.Context, request *pb.CreateAvailabilityRequest) (*pb.CreateAvailabilityResponse, error) {
	// TODO: Create availability
	print(request.Availability)
	return nil, nil
}

func (handler *AvailabilityHandler) GetAllUnavailable(ctx context.Context, request *pb.GetAllUnavailableRequest) (*pb.GetAllUnavailableResponse, error) {
	startDate := request.StartDate.AsTime()
	endDate := request.EndDate.AsTime()
	ids, err := handler.service.GetAllUnavailable(startDate,endDate)
	print(ids)
	return nil,errors.New("greska")

}

