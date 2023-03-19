package repository

// Imports
import (
	"context"
	"flight/model"

	"time"

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

func (repository *FlightRepository) SearchFlights(availableSeats int, departure time.Time, origin, destination string) ([]primitive.M, error) {
	/*
	filter := bson.M{
		"availableseats": bson.M{"$gte": availableSeats},
		"departure": bson.M{
			"$gte": departure,
			"$lt":  departure.Add(time.Hour * 23),
		},
		"origin.city":      origin,
		"destination.city": destination,
	}*/
	filter := bson.M{
		"availableseats": bson.M{"$gte": availableSeats},
	}
	
	if !departure.IsZero() {
		filter["departure"] = bson.M{
			"$gte": departure,
			"$lt":  departure.Add(time.Hour * 23),
		}
	}
	if origin != "" {
		filter["origin.city"] = origin
	}
	if destination != "" {
		filter["destination.city"] = destination
	}


	cursor, err := repository.Database.Collection("flights").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var flights []primitive.M
	for cursor.Next(context.Background()) {
		var flight bson.M
		err := cursor.Decode(&flight)
		if err != nil {
			return nil, err
		}

		price := flight["price"].(float64)
		totalPrice := price * float64(availableSeats)

		flight["totalPrice"] = totalPrice

		flights = append(flights, flight)
	}

	return flights, nil
}
