package api

import (
	pb "github.com/s-matke/abp/abp-back/common/proto/reservation_service"
	"github.com/s-matke/abp/abp-back/reservation_service/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapReservation(reservation *domain.Reservation) *pb.Reservation {
	reservationPb := &pb.Reservation{
		Id:              reservation.Id.Hex(),
		AccommodationId: reservation.AccommodationId.Hex(),
		GuestId:         reservation.GuestId,
		StartDate:       timestamppb.New(reservation.StartDate),
		EndDate:         timestamppb.New(reservation.EndDate),
		NumOfGuests:     int32(reservation.NumOfGuests),
		Price:           reservation.Price,
		Status:          mapStatus(reservation.Status),
	}

	return reservationPb
}

func mapStatus(status domain.Status) pb.Reservation_Status {
	switch status {
	case domain.BOOKED:
		return pb.Reservation_BOOKED
	}
	return pb.Reservation_CANCELLED
}
