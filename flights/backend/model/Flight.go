package model

import (
	"time"

	//"github.com/google/uuid"
	//"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	//ID             uuid.UUID `bson:"id" json:"id"`
	Departure      time.Time `bson:"departure" json:"departure"`
	Origin         Location  `bson:"origin" json:"origin"`
	Destination    Location  `bson:"destination" json:"destination"`
	Price          float64   `bson:"price" json:"price"`
	AvailableSeats int       `bson:"availableSeats" json:"availableSeats"`
}

/*func (flight *Flight) BeforeCreate(*mongo.Database) error {
	flight.ID = uuid.New()
	return nil
}*/
