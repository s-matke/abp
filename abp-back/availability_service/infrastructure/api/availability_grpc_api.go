package api

import (
	"context"
	"fmt"

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

func (handler *AvailabilityHandler) GetByAccommodation(ctx context.Context, request *pb.GetByAccommodationRequest) (*pb.GetByAccommodationResponse, error) {
	accommodation_id := request.Id

	availabilities, err := handler.service.GetByAccommodation(accommodation_id)

	if err != nil {
		return nil, err
	}

	response := &pb.GetByAccommodationResponse{
		Availabilities: []*pb.Availability{},
	}

	for _, availability := range availabilities {
		current := mapAvailability(availability)
		response.Availabilities = append(response.Availabilities, current)
	}

	return response, nil
}

func (handler *AvailabilityHandler) CreateAvailability(ctx context.Context, request *pb.CreateAvailabilityRequest) (*pb.CreateAvailabilityResponse, error) {
	availability := mapNewAvailability(request)
	err := handler.service.CreateAvailability(availability)
	if err != nil {
		return nil, err
	}

	response := &pb.CreateAvailabilityResponse{Availability: mapAvailability(availability)}

	return response, nil
}

func (handler *AvailabilityHandler) DeleteByData(ctx context.Context, request *pb.DeleteByDataRequest) (*pb.DeleteByDataResponse, error) {
	availability := mapDeleteAvailability(request)

	fmt.Println("Availability: ", availability.StartDate, "::::", availability.EndDate)

	err := handler.service.DeleteByData(availability)

	if err != nil {
		return nil, err
	}

	response := &pb.DeleteByDataResponse{
		Status: true,
	}

	return response, nil
}

func (handler *AvailabilityHandler) GetAllUnavailable(ctx context.Context, request *pb.GetAllUnavailableRequest) (*pb.GetAllUnavailableResponse, error) {
	date := mapDates(request)

	unavailableAccommodations, err := handler.service.GetAllUnavailable(date)

	if err != nil {
		return nil, err
	}

	response := &pb.GetAllUnavailableResponse{
		Availabilities: []*pb.Availability{},
	}

	for _, availability := range unavailableAccommodations {
		// fmt.Println("ID: ", availability.AccommodationId.Hex())
		current := mapAvailability(availability)
		response.Availabilities = append(response.Availabilities, current)
	}

	return response, nil

}
