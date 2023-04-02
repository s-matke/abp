package repository

import (
	"context"
	"flight/dto"
	"flight/model"
	"fmt"

	"time"

	"github.com/google/uuid"
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
	newTicket.ID = uuid.New()
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

	userId, err := uuid.Parse(ticket.IdUser)
	// objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil
	}

	fmt.Print(userId)

	err2 := repository.Database.Collection("users").FindOne(context.TODO(), bson.M{"id": userId})

	var user model.User
	err2.Decode(&user)
	user.OwnedTickets = append(user.OwnedTickets, newTicket)

	filterUser := bson.D{{Key: "id", Value: userId}}
	updateUser := bson.D{{Key: "$set", Value: bson.D{{Key: "ownedtickets", Value: user.OwnedTickets}}}}
	_, err3 := repository.Database.Collection("users").UpdateOne(context.TODO(), filterUser, updateUser)

	if err3 != nil {

		return nil
	}

	return nil
}

func (repository *TicketRepository) ShowUserTickets(idUser string) ([]dto.TicketPresenetDTO, error) {

	// objectId, err := primitive.ObjectIDFromHex(idUser)
	userId, err := uuid.Parse(idUser)
	if err != nil {
		return nil, nil
	}
	var result model.User
	err1 := repository.Database.Collection("users").FindOne(context.TODO(), bson.M{"id": userId}).Decode(&result)
	if err1 != nil {
		return nil, nil
	}
	// var tickets []model.Ticket
	tickets := result.OwnedTickets
	var ticketPresent []dto.TicketPresenetDTO

	for _, item := range tickets {
		var tPresent dto.TicketPresenetDTO
		tPresent.Destination = item.Flight.Destination
		tPresent.Origin = item.Flight.Origin
		tPresent.Price = item.Flight.Price
		tPresent.NumberOfTickets = item.NumberOfTickets
		tPresent.Departure = item.Flight.Departure.Format("2006-01-02 15:04:05")
		ticketPresent = append(ticketPresent, tPresent)
	}
	return ticketPresent, nil
}
