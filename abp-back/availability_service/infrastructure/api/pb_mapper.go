package api

import (
	"github.com/s-matke/abp/abp-back/availability_service/domain"
	pb "github.com/s-matke/abp/abp-back/common/proto/availability_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func mapNewAvailability(request *pb.CreateAvailabilityRequest) *domain.Availability {
	accId, err := primitive.ObjectIDFromHex(request.Availability.AccommodationId)

	if err != nil {
		return nil
	}

	availability := &domain.Availability{
		Id:              primitive.NewObjectID(),
		AccommodationId: accId,
		StartDate:       request.Availability.StartDate.AsTime(),
		EndDate:         request.Availability.EndDate.AsTime(),
	}

	return availability
}
