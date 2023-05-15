package api

import (
	"context"

	pb "github.com/s-matke/abp/abp-back/common/proto/reservation_service"
	"github.com/s-matke/abp/abp-back/reservation_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationHandler struct {
	pb.UnimplementedReservationServiceServer
	service *application.ReservationService
}

func NewReservationHandler(service *application.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		service: service,
	}
}

func (handler *ReservationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	reservations, err := handler.service.GetAll()

	if err != nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		Reservations: []*pb.Reservation{},
	}

	for _, reservation := range reservations {
		current := mapReservation(reservation)
		response.Reservations = append(response.Reservations, current)
	}

	return response, nil
}

func (handler *ReservationHandler) GetByAccommodation(ctx context.Context, request *pb.GetByAccommodationRequest) (*pb.GetByAccommodationResponse, error) {
	accommodation_id, err := primitive.ObjectIDFromHex(request.Id)

	if err != nil {
		return nil, err
	}

	reservations, err := handler.service.GetByAccommodation(accommodation_id)

	if err != nil {
		return nil, err
	}

	response := &pb.GetByAccommodationResponse{
		Reservations: []*pb.Reservation{},
	}

	for _, reservation := range reservations {
		current := mapReservation(reservation)
		response.Reservations = append(response.Reservations, current)
	}

	return response, nil
}

func (handler *ReservationHandler) CreateReservation(ctx context.Context, request *pb.CreateReservationRequest) (*pb.CreateReservationResponse, error) {
	// TODO
	reservation := mapNewReservation(request)

	reservation, err := handler.service.CreateReservation(reservation)

	if err != nil {
		return nil, err
	}

	response := &pb.CreateReservationResponse{
		Reservation: mapReservation(reservation),
	}

	return response, nil
}

/*
message NewReservation {
    string accommodation_id = 1;
    string guest_id = 2;
    google.protobuf.Timestamp startDate = 3;
    google.protobuf.Timestamp endDate = 4;
    int32 num_of_guests = 5;
}
*/
