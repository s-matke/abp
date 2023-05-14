package startup

import (
	"github.com/s-matke/abp/abp-back/accommodation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var accommodations = []*domain.Accommodation{
	{
		Id:     getObjectId("1"),
		HostId: "Host1_Id",
		Name:   "Naziv 1",
		Location: domain.Location{
			Address: "Adresa 1",
			City:    "City 1",
			Country: "Drzava 1",
		},
		Utilities: []domain.Utility{
			{
				Name:        "Wi-Fi",
				Description: "Free Wi-Fi",
			},
			{
				Name:        "Parking",
				Description: "Free parking spot",
			},
		},
		Images:               []string{"image1", "image2"},
		MinPeople:            1,
		MaxPeople:            4,
		AutomaticReservation: false,
	},
	{
		Id:     getObjectId("2"),
		HostId: "Host2_Id",
		Name:   "Naziv 2",
		Location: domain.Location{
			Address: "Adresa 2",
			City:    "City 2",
			Country: "Drzava 2",
		},
		Utilities: []domain.Utility{
			{
				Name:        "Wi-Fi",
				Description: "Free Wi-Fi",
			},
			{
				Name:        "Parking",
				Description: "Free parking spot",
			},
		},
		Images:               []string{"image2", "image3"},
		MinPeople:            2,
		MaxPeople:            9,
		AutomaticReservation: true,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
