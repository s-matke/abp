package api

import (
	"context"

	"github.com/s-matke/abp/abp-back/accommodation_service/application"
	pb "github.com/s-matke/abp/abp-back/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	service *application.AccommodationService
}

func NewAccommodationHandler(service *application.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		service: service,
	}
}

func (handler *AccommodationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	accommodation, err := handler.service.Get(objectId)

	if err != nil {
		return nil, err
	}

	accommodationPb := mapAccommodation(accommodation)
	response := &pb.GetResponse{
		Accommodation: accommodationPb,
	}

	return response, nil
}

func (handler *AccommodationHandler) GetByHost(ctx context.Context, request *pb.GetByHostRequest) (*pb.GetByHostResponse, error) {
	id := request.Id

	accommodations, err := handler.service.GetByHost(id)

	if err != nil {
		return nil, err
	}

	response := &pb.GetByHostResponse{
		Accommodations: []*pb.Accommodation{},
	}

	for _, accommodation := range accommodations {
		current := mapAccommodation(accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}

	return response, nil
}

func (handler *AccommodationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	accommodations, err := handler.service.GetAll()

	if err != nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		Accommodations: []*pb.Accommodation{},
	}

	for _, accommodation := range accommodations {
		current := mapAccommodation(accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}

	return response, nil
}

func (handler *AccommodationHandler) CreateAccommodation(ctx context.Context, request *pb.CreateAccommodationRequest) (*pb.CreateAccommodationResponse, error) {
	accommodation := mapNewAccommodation(request)
	err := handler.service.CreateAccommodation(accommodation)
	if err != nil {
		return nil, err
	}
	response := pb.CreateAccommodationResponse{Accommodation: mapAccommodation(accommodation)}
	return &response, nil
}
func (handler *AccommodationHandler) Search(ctx context.Context, request *pb.SearchRequest) (*pb.SearchResponse, error) {
	print("Hello world")
	accommodations, err := handler.service.GetAccommodationsBySearchCriteria(request.NumOfPeople, request.Location.City)
	//var err error
	if err != nil {
		return nil, err
	}
	response := &pb.SearchResponse{
		Accommodations: []*pb.Accommodation{},
	}

	for _, accommodation := range accommodations {
		current := mapAccommodation(accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}

	return response, nil
}
