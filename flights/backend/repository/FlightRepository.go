package repository

// Imports
import (
	"context"
	"flight/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"fmt"
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
/*
func (repository *FlightRepository) GetFlightsByAvailableSeats(availableSeats int) ([]primitive.M, error) {
	cursor, err := repository.Database.Collection("flights").Find(context.Background(), bson.M{"AvailableSeats": bson.M{"$gte": availableSeats}})
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
*/
func (repository *FlightRepository) SearchFlights(availableSeats int, departure time.Time, origin, destination string) ([]primitive.M, error) {         

	fmt.Println("Vreme- ", departure, "\nDodato 23h: ",departure.Add(time.Hour * 23))
	filter := bson.M{
		"AvailableSeats": bson.M{"$gte": availableSeats},
		"Departure": bson.M{
			"$gte": departure,
			"$lte": departure.Add(time.Hour * 23),
		},
		//"Origin.city":    origin,
		//"Destination.city": destination,
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

		price := flight["Price"].(int32)
		totalPrice := price * int32(availableSeats)

		flight["TotalPrice"] = totalPrice

		flights = append(flights, flight)
	}

	return flights, nil
}
