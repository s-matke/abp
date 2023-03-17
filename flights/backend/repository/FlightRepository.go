package repository

// Imports
import (
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"flight/model"
)

type FlightRepository struct {
	Database *mongo.Database
}

// CRUD i kompleksne funkcionalnosti
// letovi := repository.Database.Collection("letovi")

func (repository *FlightRepository) GetAllFlights() ([]model.Flight, error) {
	flightCol := repository.Database.Collection("flights")

	var flights []model.Flight
	cursor, err := flightCol.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &flights)
	if err != nil {
		return nil, err
	}
	return flights, nil
}
