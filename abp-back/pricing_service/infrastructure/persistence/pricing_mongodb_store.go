package persistence

import (
	"context"

	"github.com/s-matke/abp/abp-back/pricing_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "pricing"
	COLLECTION = "pricing"
)

type PricingMongoDBStore struct {
	pricings *mongo.Collection
}

func NewPricingMongoDBStore(client *mongo.Client) domain.PricingStore {
	pricings := client.Database(DATABASE).Collection(COLLECTION)
	return &PricingMongoDBStore{
		pricings: pricings,
	}
}

func (store *PricingMongoDBStore) Get(id primitive.ObjectID) (*domain.Pricing, error) {
	filter := bson.M{"_id": id}
	return store.FilterOne(filter)
}

func (store *PricingMongoDBStore) GetByAccommodation(id string) (*domain.Pricing, error) {
	filter := bson.M{"accommodation_id": id}
	return store.FilterOne(filter)
}

func (store *PricingMongoDBStore) FilterOne(filter interface{}) (Pricing *domain.Pricing, err error) {
	result := store.pricings.FindOne(context.TODO(), filter)
	err = result.Decode(&Pricing)
	return
}

func (store *PricingMongoDBStore) DeleteAll() {
	store.pricings.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *PricingMongoDBStore) GetAll() ([]*domain.Pricing, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *PricingMongoDBStore) filter(filter interface{}) ([]*domain.Pricing, error) {
	cursor, err := store.pricings.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (accommodations []*domain.Pricing, err error) {
	for cursor.Next(context.TODO()) {
		var Accommodation domain.Pricing
		err = cursor.Decode(&Accommodation)
		if err != nil {
			return
		}
		accommodations = append(accommodations, &Accommodation)
	}
	err = cursor.Err()
	return
}

func (store *PricingMongoDBStore) Insert(Accommodation *domain.Pricing) error {
	result, err := store.pricings.InsertOne(context.TODO(), Accommodation)
	if err != nil {
		return err
	}
	Accommodation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}
