package persistence

import (
	"context"
	"fmt"

	"github.com/mihajlo-ra92/XML/notifications_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "notifications"
	COLLECTION = "notifications"
)

type NotificationsMongoDBStore struct {
	Notifications *mongo.Collection
}

func NewNotificationsMongoDBStore(client *mongo.Client) domain.NotificationsStore {
	Notifications := client.Database(DATABASE).Collection(COLLECTION)
	return &NotificationsMongoDBStore{
		Notifications: Notifications,
	}
}

func (store *NotificationsMongoDBStore) Get(id primitive.ObjectID) (*domain.Notification, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *NotificationsMongoDBStore) GetByUser(userId string) (*domain.Notification, error) {
	filter := bson.M{"user_id": userId}
	return store.filterOne(filter)
}

func (store *NotificationsMongoDBStore) GetAll() ([]*domain.Notification, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *NotificationsMongoDBStore) Insert(Notification *domain.Notification) error {
	result, err := store.Notifications.InsertOne(context.TODO(), Notification)
	if err != nil {
		return err
	}
	Notification.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *NotificationsMongoDBStore) filter(filter interface{}) ([]*domain.Notification, error) {
	cursor, err := store.Notifications.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *NotificationsMongoDBStore) filterOne(filter interface{}) (Notification *domain.Notification, err error) {
	result := store.Notifications.FindOne(context.TODO(), filter)
	err = result.Decode(&Notification)
	return
}

func (store *NotificationsMongoDBStore) DeleteAll() {
	store.Notifications.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *NotificationsMongoDBStore) Delete(Notification *domain.Notification) error {
	filter := bson.M{"_id": Notification.Id}
	result, err := store.Notifications.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with ID %s", Notification.Id)
	}
	return nil
}

func decode(cursor *mongo.Cursor) (Notifications []*domain.Notification, err error) {
	for cursor.Next(context.TODO()) {
		var Notification domain.Notification
		err = cursor.Decode(&Notification)
		if err != nil {
			return
		}
		Notifications = append(Notifications, &Notification)
	}
	err = cursor.Err()
	return
}
