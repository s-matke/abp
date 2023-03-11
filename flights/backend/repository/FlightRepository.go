package repository

// Imports
import (
	"go.mongodb.org/mongo-driver/mongo"
)

type FlightRepository struct {
	Database *mongo.Database
}

// CRUD i kompleksne funkcionalnosti
// letovi := repository.Database.Collection("letovi")
