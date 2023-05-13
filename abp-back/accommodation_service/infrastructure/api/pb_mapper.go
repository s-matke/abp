package api

import (
	"github.com/s-matke/abp/abp-back/accommodation_service/domain"
	pb "github.com/s-matke/abp/abp-back/common/proto/accommodation_service"
)

func mapAccommodation(accommodation *domain.Accommodation) *pb.Accommodation {
	accommodationPb := &pb.Accommodation{
		Id:        accommodation.Id.Hex(),
		HostId:    accommodation.HostId,
		Name:      accommodation.Name,
		Images:    accommodation.Images,
		MinPeople: int32(accommodation.MinPeople),
		MaxPeople: int32(accommodation.MaxPeople),
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
