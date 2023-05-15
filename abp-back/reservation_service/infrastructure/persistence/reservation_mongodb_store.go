package persistence

import (
	"context"
	"fmt"

	"github.com/s-matke/abp/abp-back/reservation_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE    = "reservation"
	COLLLECTION = "reservation"
)

type ReservationMongoDBStore struct {
	reservations *mongo.Collection
}

func NewReservationMongoDBStore(client *mongo.Client) domain.ReservationStore {
	reservations := client.Database(DATABASE).Collection(COLLLECTION)
	return &ReservationMongoDBStore{
		reservations: reservations,
	}
}

func (store *ReservationMongoDBStore) GetAll() ([]*domain.Reservation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *ReservationMongoDBStore) GetByAccommodation(id primitive.ObjectID) ([]*domain.Reservation, error) {
	filter := bson.M{"accommodation_id": id}
	return store.filter(filter)
}

func (store *ReservationMongoDBStore) GetByGuest(id string) ([]*domain.Reservation, error) {
	filter := bson.M{"guest_id": id, "status": bson.M{"$ne": domain.CANCELLED}}
	return store.filter(filter)
}

func (store *ReservationMongoDBStore) GetCancelledAmount(id string) int32 {
	filter := bson.M{"guest_id": id, "status": domain.CANCELLED}
	count, err := store.reservations.CountDocuments(context.TODO(), filter)
	if err != nil {
		fmt.Println("Nije pronasao ni jednog user-a sa CANCELLED.")
		return 0
	}
	return int32(count)
}

func (store *ReservationMongoDBStore) GetAllPendingByAccommodation(id primitive.ObjectID) ([]*domain.Reservation, error) {
	filter := bson.M{"accommodation_id": id, "status": domain.PENDING}
	return store.filter(filter)
}

func (store *ReservationMongoDBStore) ConfirmReservation(id primitive.ObjectID) (*domain.Reservation, error) {
	filter := bson.M{"_id": id}

	// Define the update
	update := bson.M{"$set": bson.M{"status": domain.BOOKED}}

	_, err := store.reservations.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return nil, err
	}

	fmt.Println("UpdateOne done")

	return store.filterOne(filter)
}

func (store *ReservationMongoDBStore) filterOne(filter interface{}) (reservation *domain.Reservation, err error) {
	result := store.reservations.FindOne(context.TODO(), filter)
	err = result.Decode(&reservation)
	return
}

func (store *ReservationMongoDBStore) Insert(reservation *domain.Reservation) error {
	result, err := store.reservations.InsertOne(context.TODO(), reservation)

	if err != nil {
		return err
	}

	reservation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *ReservationMongoDBStore) DeleteAll() {
	store.reservations.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *ReservationMongoDBStore) DeleteByIds(ids []primitive.ObjectID) error {
	filter := bson.M{
		"_id": bson.M{"$in": ids},
	}

	_, err := store.reservations.DeleteMany(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}

func (store *ReservationMongoDBStore) filter(filter interface{}) ([]*domain.Reservation, error) {
	cursor, err := store.reservations.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (reservations []*domain.Reservation, err error) {
	for cursor.Next(context.TODO()) {
		var reservation domain.Reservation
		err = cursor.Decode(&reservation)
		if err != nil {
			return
		}
		reservations = append(reservations, &reservation)
	}
	err = cursor.Err()
	return
}
