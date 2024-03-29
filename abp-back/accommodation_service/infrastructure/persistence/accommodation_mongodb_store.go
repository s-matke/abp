package persistence

import (
	"context"
	"fmt"

	"github.com/s-matke/abp/abp-back/accommodation_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "accommodation"
	COLLECTION = "accommodation"
)

type AccommodationMongoDBStore struct {
	accommodations *mongo.Collection
}

func NewAccommodationMongoDBStore(client *mongo.Client) domain.AccommodationStore {
	accommodations := client.Database(DATABASE).Collection(COLLECTION)
	return &AccommodationMongoDBStore{
		accommodations: accommodations,
	}
}

func (store *AccommodationMongoDBStore) Get(id primitive.ObjectID) (*domain.Accommodation, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AccommodationMongoDBStore) GetAll() ([]*domain.Accommodation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AccommodationMongoDBStore) Insert(Accommodation *domain.Accommodation) error {
	result, err := store.accommodations.InsertOne(context.TODO(), Accommodation)
	if err != nil {
		return err
	}
	Accommodation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *AccommodationMongoDBStore) DeleteAll() {
	store.accommodations.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *AccommodationMongoDBStore) filterOne(filter interface{}) (Accommodation *domain.Accommodation, err error) {
	result := store.accommodations.FindOne(context.TODO(), filter)
	err = result.Decode(&Accommodation)
	return
}

func (store *AccommodationMongoDBStore) filter(filter interface{}) ([]*domain.Accommodation, error) {
	cursor, err := store.accommodations.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (accommodations []*domain.Accommodation, err error) {
	for cursor.Next(context.TODO()) {
		var Accommodation domain.Accommodation
		err = cursor.Decode(&Accommodation)
		if err != nil {
			return
		}
		accommodations = append(accommodations, &Accommodation)
	}
	err = cursor.Err()
	return
}

func (store *AccommodationMongoDBStore) GetByHost(id string) ([]*domain.Accommodation, error) {
	filter := bson.M{"host_id": id}
	print(store.filter(filter))
	return store.filter(filter)
}
func (store *AccommodationMongoDBStore) SearchAccommodation(availableSeats int32, destination string, ids []primitive.ObjectID) ([]*domain.Accommodation, error) {

	baseFilter := bson.M{
		"minPeople": bson.M{"$lte": availableSeats},
		"maxPeople": bson.M{"$gte": availableSeats},
	}

	additionalConditions := bson.M{}

	if len(ids) != 0 {
		fmt.Println("NIJE == 0")
		additionalConditions["_id"] = bson.M{"$nin": ids}
	}

	cityFilter := bson.M{}
	if destination != "" {
		cityFilter["location.city"] = destination
	}

	filter := bson.M{
		"$and": []bson.M{baseFilter, additionalConditions, cityFilter},
	}

	cursor, err := store.accommodations.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var accommodations []*domain.Accommodation
	for cursor.Next(context.Background()) {
		var accommodation *domain.Accommodation
		err := cursor.Decode(&accommodation)
		if err != nil {
			return nil, err
		}

		accommodations = append(accommodations, accommodation)
	}

	return accommodations, nil
}
