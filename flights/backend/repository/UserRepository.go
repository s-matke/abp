package repository

import (
	"context"
	"flight/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Database *mongo.Database
}

func (repository *UserRepository) CreateUser(user *model.User) error {
	dbResult, err := repository.Database.Collection("users").InsertOne(context.TODO(), &user)
	if err != nil {
		return err
	}

	println("Successfully added user: ", dbResult.InsertedID.(primitive.ObjectID).Hex())
	return nil
}

func (repository *UserRepository) FindByEmail(email string) (bson.D, error) {
	userCol := repository.Database.Collection("users")

	var user bson.D
	err := userCol.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}
	return user, nil
}
