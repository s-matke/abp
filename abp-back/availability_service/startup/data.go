package startup

import (
	"time"

	"github.com/s-matke/abp/abp-back/availability_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var availabilities = []*domain.Availability{
	{
		Id:              getObjectId("1"),
		AccommodationId: getObjectId("64621a0d4a9b18ce6897930e"),
		StartDate:       time.Now(),
		EndDate:         time.Now().AddDate(0, 0, 10),
	},
	{
		Id:              getObjectId("3"),
		AccommodationId: getObjectId("64621a0d4a9b18ce6897930d"),
		StartDate:       time.Now(),
		EndDate:         time.Now().AddDate(0, 0, 2),
	},
	{
		Id:              getObjectId("4"),
		AccommodationId: getObjectId("64621a0d4a9b18ce6897930d"),
		StartDate:       time.Now(),
		EndDate:         time.Now().AddDate(0, 0, 5),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
