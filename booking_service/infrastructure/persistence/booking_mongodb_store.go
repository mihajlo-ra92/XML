package persistence

import (
	"context"

	"github.com/mihajlo-ra92/XML/booking_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "booking"
	COLLECTION = "booking"
)

type BookingMongoDBStore struct {
	bookings *mongo.Collection
}

func NewBookingMongoDBStore(client *mongo.Client) domain.BookingStore {
	bookings := client.Database(DATABASE).Collection(COLLECTION)
	return &BookingMongoDBStore{
		bookings: bookings,
	}
}

func (store *BookingMongoDBStore) Get(id primitive.ObjectID) (*domain.Booking, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *BookingMongoDBStore) GetAll() ([]*domain.Booking, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *BookingMongoDBStore) Insert(Booking *domain.Booking) error {
	result, err := store.bookings.InsertOne(context.TODO(), Booking)
	if err != nil {
		return err
	}
	Booking.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *BookingMongoDBStore) DeleteAll() {
	store.bookings.DeleteMany(context.TODO(), bson.D{{}})
}

// func (store *UserMongoDBStore) Update

func (store *BookingMongoDBStore) filter(filter interface{}) ([]*domain.Booking, error) {
	cursor, err := store.bookings.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *BookingMongoDBStore) filterOne(filter interface{}) (Booking *domain.Booking, err error) {
	result := store.bookings.FindOne(context.TODO(), filter)
	err = result.Decode(&Booking)
	return
}

func decode(cursor *mongo.Cursor) (bookings []*domain.Booking, err error) {
	for cursor.Next(context.TODO()) {
		var booking domain.Booking
		err = cursor.Decode(&booking)
		if err != nil {
			return
		}
		bookings = append(bookings, &booking)
	}
	err = cursor.Err()
	return
}
