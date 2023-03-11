package repository

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct {
	Database *mongo.Database
}
