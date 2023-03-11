package repository

import "go.mongodb.org/mongo-driver/mongo"

type TicketRepository struct {
	Database *mongo.Database
}
