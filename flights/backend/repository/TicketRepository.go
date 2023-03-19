package repository

import (
	"context"
	"flight/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TicketRepository struct {
	Database *mongo.Database
}

func (repository *TicketRepository) BuyTicket(ticket *model.DTOTicket) error {

	var newTicket model.Ticket
	newTicket.Flight = ticket.Flight
	newTicket.IssuedAt = time.Now()
	newTicket.NumberOfTickets = ticket.NumberOfTickets
	_, err := repository.Database.Collection("tickets").InsertOne(context.TODO(), &newTicket)

	if err != nil {
		return err
	}

	flightId, _ := primitive.ObjectIDFromHex(ticket.IdFlight)
	filterFlight := bson.D{{Key: "_id", Value: flightId}}
	updateFlight := bson.D{{Key: "$set", Value: bson.D{{Key: "availableseats", Value: ticket.Flight.AvailableSeats}}}}
	_, err1 := repository.Database.Collection("flights").UpdateOne(context.TODO(), filterFlight, updateFlight)

	if err1 != nil {
		return err
	}

	userId := ticket.IdUser
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil
	}

	err2 := repository.Database.Collection("users").FindOne(context.TODO(), bson.M{"_id": objectId})

	var user model.User
	err2.Decode(&user)
	user.OwnedTickets = append(user.OwnedTickets, newTicket)

	filterUser := bson.D{{Key: "_id", Value: objectId}}
	updateUser := bson.D{{Key: "$set", Value: bson.D{{Key: "ownedtickets", Value: user.OwnedTickets}}}}
	_, err3 := repository.Database.Collection("users").UpdateOne(context.TODO(), filterUser, updateUser)

	if err3 != nil {
		return nil
	}

	return nil
}
