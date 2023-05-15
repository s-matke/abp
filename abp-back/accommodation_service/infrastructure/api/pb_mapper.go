package api

import (
	"github.com/s-matke/abp/abp-back/accommodation_service/domain"
	pb "github.com/s-matke/abp/abp-back/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapAccommodation(accommodation *domain.Accommodation) *pb.Accommodation {
	accommodationPb := &pb.Accommodation{
		Id:                   accommodation.Id.Hex(),
		HostId:               accommodation.HostId,
		Name:                 accommodation.Name,
		Images:               accommodation.Images,
		MinPeople:            int32(accommodation.MinPeople),
		MaxPeople:            int32(accommodation.MaxPeople),
		AutomaticReservation: accommodation.AutomaticReservation,
	}

	accommodationPb.Location = &pb.Location{
		Address: accommodation.Location.Address,
		City:    accommodation.Location.City,
		Country: accommodation.Location.Country,
	}

	for _, utility := range accommodation.Utilities {
		accommodationPb.Utilities = append(accommodationPb.Utilities, &pb.Utility{
			Name:        utility.Name,
			Description: utility.Description,
		})
	}

	return accommodationPb
}

func mapNewAccommodation(request *pb.CreateAccommodationRequest) *domain.Accommodation {
	accommodation := &domain.Accommodation{
		Id:        primitive.NewObjectID(),
		HostId:    request.Accommodation.HostId,
		Name:      request.Accommodation.Name,
		Images:    request.Accommodation.Images,
		MinPeople: uint32(request.Accommodation.MinPeople),
		MaxPeople: uint32(request.Accommodation.MaxPeople),
		Utilities: make([]domain.Utility, 0),
	}

	location := &domain.Location{
		Address: request.Accommodation.Location.Address,
		City:    request.Accommodation.Location.City,
		Country: request.Accommodation.Location.Country,
	}

	for _, utilitiesPb := range request.Accommodation.Utilities {
		utility := domain.Utility{
			Name:        utilitiesPb.Name,
			Description: utilitiesPb.Description,
		}
		accommodation.Utilities = append(accommodation.Utilities, utility)
	}

	accommodation.Location = *location

	return accommodation
}
