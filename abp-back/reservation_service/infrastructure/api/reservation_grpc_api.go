package api

import (
	"context"
	"fmt"

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

func (handler *ReservationHandler) GetAllPendingByAccommodation(ctx context.Context, request *pb.GetAllPendingByAccommodationRequest) (*pb.GetAllPendingByAccommodationResponse, error) {
	accommodation_id, err := primitive.ObjectIDFromHex(request.Id)

	if err != nil {
		return nil, err
	}

	fmt.Println(accommodation_id)
	fmt.Println(request.Id)

	reservations, err := handler.service.GetAllPendingByAccommodation(accommodation_id)

	if err != nil {
		return nil, err
	}

	response := &pb.GetAllPendingByAccommodationResponse{
		Reservations: []*pb.PendingReservation{},
	}

	for _, reservation := range reservations {
		num_of_cancelled := handler.service.GetCancelledAmount(reservation.GuestId)
		current := mapPendingReservation(reservation, num_of_cancelled)
		response.Reservations = append(response.Reservations, current)
	}

	fmt.Println("Prosao kroz for loop: ", response)

	return response, nil
}

func (handler *ReservationHandler) ConfirmReservation(ctx context.Context, request *pb.ConfirmReservationRequest) (*pb.ConfirmReservationResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)

	if err != nil {
		return nil, err
	}
	fmt.Println("ID: ", id)
	reservations, err := handler.service.ConfirmReservation(id)

	if err != nil {
		return nil, err
	}

	response := &pb.ConfirmReservationResponse{
		Reservations: []*pb.Reservation{},
	}

	for _, reservation := range reservations {
		current := mapReservation(reservation)
		response.Reservations = append(response.Reservations, current)
	}

	return response, nil
}

func (handler *ReservationHandler) GetByGuest(ctx context.Context, request *pb.GetByGuestRequest) (*pb.GetByGuestResponse, error) {
	id := request.Id

	reservations, err := handler.service.GetByGuest(id)

	if err != nil {
		return nil, err
	}

	response := &pb.GetByGuestResponse{
		Reservations: []*pb.Reservation{},
	}

	for _, reservation := range reservations {
		current := mapReservation(reservation)
		response.Reservations = append(response.Reservations, current)
	}

	return response, nil
}
