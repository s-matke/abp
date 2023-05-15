package startup

import (
	"time"

	"github.com/s-matke/abp/abp-back/reservation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var reservations = []*domain.Reservation{
	{
		Id:              getObjectId("1"),
		AccommodationId: getObjectId("1a"),
		GuestId:         "123",
		StartDate:       time.Now(),
		EndDate:         time.Now().AddDate(0, 0, 3),
		NumOfGuests:     3,
		Price:           150,
		Status:          domain.PENDING,
	},
	{
		Id:              getObjectId("2"),
		AccommodationId: getObjectId("1b"),
		GuestId:         "123",
		StartDate:       time.Now(),
		EndDate:         time.Now().AddDate(0, 0, 1),
		NumOfGuests:     2,
		Price:           100,
		Status:          domain.BOOKED,
	},
	{
		Id:              getObjectId("3"),
		AccommodationId: getObjectId("1c"),
		GuestId:         "124",
		StartDate:       time.Now(),
		EndDate:         time.Now().AddDate(0, 0, 10),
		NumOfGuests:     5,
		Price:           300,
		Status:          domain.BOOKED,
	},
	{
		Id:              getObjectId("4"),
		AccommodationId: getObjectId("1ba"),
		GuestId:         "123",
		StartDate:       time.Now(),
		EndDate:         time.Now().AddDate(0, 0, 3),
		NumOfGuests:     3,
		Price:           150,
		Status:          domain.CANCELLED,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
