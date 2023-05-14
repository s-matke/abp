package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Availability struct {
	Id              primitive.ObjectID `bson:"_id"`
	AccommodationId primitive.ObjectID `bson:"accommodation_id"`
	StartDate       time.Time          `bson:"start_date"`
	EndDate         time.Time          `bson:"end_date"`
}
