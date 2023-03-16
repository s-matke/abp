package repository

// Imports
import (
	"context"
	"flight/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FlightRepository struct {
	Database *mongo.Database
}

func (repository *FlightRepository) Create(flight *model.Flight) error {
	dbResult, err := repository.Database.Collection("flights").InsertOne(context.TODO(), &flight)

	if err != nil {
		return err
	}

	println("Successfully added flight: ", dbResult.InsertedID.(primitive.ObjectID).Hex())
	return nil
}

func (repository *FlightRepository) GetAllFlights() ([]primitive.M, error) {
	cursor, err := repository.Database.Collection("flights").Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	var flights []primitive.M
	for cursor.Next(context.Background()) {
		var flight bson.M
		err := cursor.Decode(&flight)
		if err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}
	defer cursor.Close(context.Background())
	return flights, nil
}
