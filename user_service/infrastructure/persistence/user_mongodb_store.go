package persistence

import (
	"context"
	"fmt"

	"github.com/mihajlo-ra92/XML/user_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE = "user"
	COLLECTION = "user"
)

type UserMongoDBStore struct {
	users * mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) Get (id primitive.ObjectID) (*domain.User, error){
	filter := bson.M{"_id":id}
	return store.filterOne(filter)
}
func (store *UserMongoDBStore) GetByUsername(username string) (*domain.User, error){
	fmt.Println("in GetByUsername")
	filter := bson.M{"username":username}
	fmt.Print("filter: ")
	fmt.Println(filter)
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetByEmail(email string) (*domain.User, error){
	fmt.Println("in GetByEmail")
	filter := bson.M{"email":email}
	fmt.Print("filter: ")
	fmt.Println(filter)
	return store.filterOne(filter)
}



func (store *UserMongoDBStore) GetByLoginData(username string, password string) (*domain.User, error){
	fmt.Println("in GetByLoginData")
	filter := bson.M{"username":username}
	fmt.Print("filter: ")
	fmt.Println(filter)
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetAll() ([]*domain.User, error){
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserMongoDBStore) Insert(User *domain.User) error {
	result, err := store.users.InsertOne(context.TODO(), User)
	if err != nil {
		return err
	}
	User.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}
func (store *UserMongoDBStore) Update(user *domain.User) error {
	fmt.Print("user in mongodb_store: ")
	fmt.Println(user)
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": bson.M{
		"user_type":     user.UserType,
		"username":    user.Username,
		"password": user.Password,
		"email": user.Email,
		"first_name": user.FirstName,
		"last_name": user.LastName,
		"address": user.Address,
	}}
	updateResult, err := store.users.UpdateByID(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Print("updateResult: ")
	fmt.Print(updateResult)
	return nil
}

func (store *UserMongoDBStore) Delete(user *domain.User) error{
	filter := bson.M{"_id": user.Id}
	result, err := store.users.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with ID %s", user.Id)
	}
	return nil
}


func (store *UserMongoDBStore) DeleteAll() {
	store.users.DeleteMany(context.TODO(), bson.D{{}})
}

// func (store *UserMongoDBStore) Update

func (store *UserMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err :=  store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserMongoDBStore) filterOne(filter interface{}) (User *domain.User, err error){
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}

func decode(cursor *mongo.Cursor) (users []*domain.User, err error){
	for cursor.Next(context.TODO()){
		var user domain.User
		err = cursor.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cursor.Err()
	return
}