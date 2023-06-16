package persistence

import (
	"context"
	"github.com/mihajlo-ra92/XML/accommodation_service/domain"

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

func (store *AccommodationMongoDBStore) GetByHostId(hostId string) ([]*domain.Accommodation, error) {
	filter := bson.M{"host_id": hostId}
	return store.filter(filter)
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

func (store *AccommodationMongoDBStore) DeleteByHostId(hostId string) error{
	filter := bson.M{"host_id": hostId}
	_, err := store.accommodations.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *AccommodationMongoDBStore) filter(filter interface{}) ([]*domain.Accommodation, error) {
	cursor, err := store.accommodations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AccommodationMongoDBStore) filterOne(filter interface{}) (Accommodation *domain.Accommodation, err error) {
	result := store.accommodations.FindOne(context.TODO(), filter)
	err = result.Decode(&Accommodation)
	return
}

func decode(cursor *mongo.Cursor) (accommodations []*domain.Accommodation, err error) {
	for cursor.Next(context.TODO()) {
		var accommodation domain.Accommodation
		err = cursor.Decode(&accommodation)
		if err != nil {
			return
		}
		accommodations = append(accommodations, &accommodation)
	}
	err = cursor.Err()
	return
}

func (store *AccommodationMongoDBStore) Search(location string, numberGuests uint32) ([]*domain.Accommodation, error) {
	regex := primitive.Regex{Pattern: location, Options: "i"}

	filter := bson.M{
		"location":   regex,
		"min_guests": bson.M{"$lte": numberGuests},
		"max_guests": bson.M{"$gte": numberGuests},
	}
	cursor, err := store.accommodations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AccommodationMongoDBStore) SearchWithFilter(location string, numberGuests uint32, minPrice uint32, maxPrice uint32, benefits []string) ([]*domain.Accommodation, error) {
	regex := primitive.Regex{Pattern: location, Options: "i"}
	filter := bson.M{
		"location":   regex,
		"min_guests": bson.M{"$lte": numberGuests},
		"max_guests": bson.M{"$gte": numberGuests},
		"price":      bson.M{"$lte": maxPrice, "$gte": minPrice},
	}

	if len(benefits) > 0 {
		filter["benefits"] = bson.M{"$all": benefits}
	}

	cursor, err := store.accommodations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}

	return decode(cursor)
}