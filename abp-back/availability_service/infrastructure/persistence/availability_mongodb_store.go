package persistence

import (
	"context"

	"github.com/s-matke/abp/abp-back/availability_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "availability"
	COLLECTION = "availability"
)

type AvailabilityMongoDBStore struct {
	availabilities *mongo.Collection
}

func NewAvailabilityMongoDBStore(client *mongo.Client) domain.AvailabilityStore {
	availabilities := client.Database(DATABASE).Collection(COLLECTION)
	return &AvailabilityMongoDBStore{
		availabilities: availabilities,
	}
}

func (store *AvailabilityMongoDBStore) GetAll() ([]*domain.Availability, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AvailabilityMongoDBStore) GetByAccommodation(id string) ([]*domain.Availability, error) {
	accId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"accommodation_id": accId}
	return store.filter(filter)
}

func (store *AvailabilityMongoDBStore) DeleteByData(availability *domain.Availability) error {
	filter := bson.M{
		"accommodation_id": availability.AccommodationId,
		"startDate":        availability.StartDate,
		"endDate":          availability.EndDate,
	}
	_, err := store.availabilities.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *AvailabilityMongoDBStore) Insert(availability *domain.Availability) error {
	result, err := store.availabilities.InsertOne(context.TODO(), availability)
	if err != nil {
		return err
	}
	availability.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *AvailabilityMongoDBStore) DeleteAll() {
	store.availabilities.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *AvailabilityMongoDBStore) GetAllUnavailable(availability *domain.Availability) ([]*domain.Availability, error) {
	filter := bson.M{
		"$or": bson.A{
			bson.M{
				"startTime": bson.M{"$lt": availability.EndDate},
				"endTime":   bson.M{"$gt": availability.StartDate},
			},
		},
	}
	return store.filter(filter)
}

func (store *AvailabilityMongoDBStore) filter(filter interface{}) ([]*domain.Availability, error) {
	cursor, err := store.availabilities.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (availabilities []*domain.Availability, err error) {
	for cursor.Next(context.TODO()) {
		var availability domain.Availability
		err = cursor.Decode(&availability)
		if err != nil {
			return
		}
		availabilities = append(availabilities, &availability)
	}
	err = cursor.Err()
	return
}
