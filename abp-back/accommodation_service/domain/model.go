package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Accommodation struct {
	Id        primitive.ObjectID `bson:"_id"`
	HostId    string             `bson:"host_id"`
	Name      string             `bson:"name"`
	Location  Location           `bson:"location"`
	Utilities []Utility          `bson:"utilities"`
	Images    []string           `bson:"images"`
	MinPeople uint32             `bson:"minPeople"`
	MaxPeople uint32             `bson:"maxPeople"`
	// Price     float64            `bson:"price"`
	// Status    AvailableStatus
}

type Location struct {
	Address string `bson:"address"`
	City    string `bson:"city"`
	Country string `bson:"country"`
}

type Utility struct {
	Name        string `bson:"name"`
	Description string `bson:"description"`
}
