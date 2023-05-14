package api

import (
	"github.com/s-matke/abp/abp-back/availability_service/domain"
	pb "github.com/s-matke/abp/abp-back/common/proto/availability_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapAvailability(availability *domain.Availability) *pb.Availability {
	availabilityPb := &pb.Availability{
		Id:              availability.Id.Hex(),
		AccommodationId: availability.AccommodationId.Hex(),
		StartDate:       timestamppb.New(availability.StartDate),
		EndDate:         timestamppb.New(availability.EndDate),
	}

	return availabilityPb
}
