package startup

import (
	"github.com/s-matke/abp/abp-back/accommodation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var accommodations = []*domain.Accommodation{
	{
		Id:     getObjectId("64621a0d4a9b18ce6897930e"),
		HostId: "Host1_Id",
		Name:   "111111111111111111",
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
		AutomaticReservation: true,
	},
	{
		Id:     getObjectId("64621a0d4a9b18ce6897930d"),
		HostId: "Host2_Id",
		Name:   "22222222222222",
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
