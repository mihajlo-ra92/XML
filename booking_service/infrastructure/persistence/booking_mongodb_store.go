package persistence

import (
	"context"
	"fmt"
	"time"

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

func (store *BookingMongoDBStore) Delete(Booking *domain.Booking) error {
	filter := bson.M{"_id": Booking.Id}
	result, err := store.bookings.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with ID %s", Booking.Id)
	}
	return nil
}

func (store *BookingMongoDBStore) DeleteByGuestId(guestId string) error {
	filter := bson.M{"guest_id": guestId}
	_, err := store.bookings.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *BookingMongoDBStore) DeleteByAccommodationId(accommodationId string) error {
	filter := bson.M{"accommodation_id": accommodationId}
	_, err := store.bookings.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *BookingMongoDBStore) DeleteAll() {
	store.bookings.DeleteMany(context.TODO(), bson.D{{}})
}

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

func (store *BookingMongoDBStore) GetByAccomodationIdandDataRange(accommodationId string, startDate time.Time, endDate time.Time) ([]*domain.Booking, error) {
	filter := bson.M{
		"accommodation_id": accommodationId,
		"start_date": bson.M{
			"$lt": endDate,
		},
		"end_date": bson.M{
			"$gt": startDate,
		},
	}
	cursor, err := store.bookings.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *BookingMongoDBStore) GetAllByUser(guestId string, bookingType domain.BookingType) ([]*domain.Booking, error) {
	filter := bson.M{
		"guest_id":     guestId,
		"booking_type": bookingType,
	}
	cursor, err := store.bookings.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *BookingMongoDBStore) GetByAccommodationId(accommodationId string) ([]*domain.Booking, error) {
	filter := bson.M{
		"accommodation_id":     accommodationId,
	}
	cursor, err := store.bookings.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *BookingMongoDBStore) Update(booking *domain.Booking) (*domain.Booking, error) {
	fmt.Print("booking in mongodb_store: ")
	fmt.Println(booking)
	filter := bson.M{"_id": booking.Id}
	update := bson.M{"$set": bson.M{
		"_id":              booking.Id,
		"accommodation_id": booking.AccommodationId,
		"guest_id":         booking.GuestId,
		"price":            booking.Price,
		"price_type":       booking.PriceType,
		"number_of_guests": booking.NumberOfGuests,
		"booking_type":     booking.BookingType,
		"start_date":       booking.StartDate,
		"end_date":         booking.EndDate,
	}}
	updateResult, err := store.bookings.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	fmt.Print("updateResult: ")
	fmt.Print(updateResult)
	return booking, nil
}
