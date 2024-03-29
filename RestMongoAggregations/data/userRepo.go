package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	// NoSQL: module containing Mongo api client
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NoSQL: ProductRepo struct encapsulating Mongo api client
type PatientRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

type UserRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

// NoSQL: Constructor which reads db configuration from environment
func NewUserRepo(ctx context.Context, logger *log.Logger) (*UserRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &UserRepo{
		cli:    client,
		logger: logger,
	}, nil
}

func (ur *UserRepo) DisconnectUserRepo(ctx context.Context) error {
	err := ur.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ur.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		ur.logger.Println(err)
	}

	databases, err := ur.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		ur.logger.Println(err)
	}
	fmt.Println(databases)
}

func (ur *UserRepo) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbName := "mongoDemo"
	usersCollectionName := "users"
	userDatabase := ur.cli.Database(dbName)

	test := os.Getenv("TEST")

	if test == "YES" {
		usersCollection := userDatabase.Collection(usersCollectionName + "_test")
		usersCollection.Drop(ctx)
		return nil
	}
	return nil
}

func (ur *UserRepo) LoginUser(username string, password string) (*User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ur.logger.Println("Username: " + username)
	ur.logger.Println("Password: " + password)

	usersCollectoin := ur.getCollection()
	var user User
	err := usersCollectoin.FindOne(ctx, bson.M{"username": username, "password": password}).Decode(&user)
	if err != nil {
		ur.logger.Println(err)
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepo) GetAll() (Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := ur.getCollection()

	var users Users
	usersCursor, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		ur.logger.Println(err)
		return nil, err
	}
	if err = usersCursor.All(ctx, &users); err != nil {
		ur.logger.Println(err)
		return nil, err
	}
	return users, nil
}

func (pr *PatientRepo) GetById(id string) (*Patient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	patientsCollection := pr.getCollection()

	var patient Patient
	objID, _ := primitive.ObjectIDFromHex(id)
	err := patientsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&patient)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return &patient, nil
}

func (ur *UserRepo) GetByUsername(username string) (Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollectoin := ur.getCollection()
	var users Users
	usersCursor, err := usersCollectoin.Find(ctx, bson.M{"username": username})
	if err != nil {
		ur.logger.Println(err)
		return nil, err
	}
	if err = usersCursor.All(ctx, &users); err != nil {
		ur.logger.Println(err)
		return nil, err
	}
	return users, nil
}

func (ur *UserRepo) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := ur.getCollection()

	result, err := usersCollection.InsertOne(ctx, &user)
	if err != nil {
		ur.logger.Println(err)
		return err
	}
	ur.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (ur *UserRepo) Update(id string, user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := ur.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"username": user.Username,
		"password": user.Password,
	}}
	result, err := usersCollection.UpdateOne(ctx, filter, update)
	ur.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	ur.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		ur.logger.Println(err)
		return err
	}
	return nil
}

func (pr *PatientRepo) AddPhoneNumber(id string, phoneNumber string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$push": bson.M{
		"phoneNumbers": phoneNumber,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (ur *UserRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := ur.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := usersCollection.DeleteOne(ctx, filter)
	if err != nil {
		ur.logger.Println(err)
		return err
	}
	ur.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

func (pr *PatientRepo) AddAnamnesis(id string, anamnesis *Anamnesis) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$push": bson.M{
		"anamnesis": anamnesis,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (pr *PatientRepo) AddTherapy(id string, therapy *Therapy) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$push": bson.M{
		"therapy": therapy,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (pr *PatientRepo) UpdateAddress(id string, address *Address) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$set": bson.M{
		"address": address,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (pr *PatientRepo) ChangePhone(id string, index int, phoneNumber string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	// What happens if set value for index=10, but we only have 3 phone numbers?
	// -> Every value in between will be set to an empty string
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$set": bson.M{
		"phoneNumbers." + strconv.Itoa(index): phoneNumber,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

// BONUS
func (pr *PatientRepo) Receipt(id string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	matchStage := bson.D{{"$match", bson.D{{"_id", objID}}}}
	sumStage := bson.D{{"$addFields",
		bson.D{{"total", bson.D{{"$add",
			bson.D{{"$sum", "$therapy.price"}},
		}},
		}},
	}}
	projectStage := bson.D{{"$project",
		bson.D{{"total", 1}},
	}}

	cursor, err := patientsCollection.Aggregate(ctx, mongo.Pipeline{matchStage, sumStage, projectStage})
	if err != nil {
		pr.logger.Println(err)
		return -1, err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		pr.logger.Println(err)
		return -1, err
	}
	for _, result := range results {
		pr.logger.Println(result)
		return result["total"].(float64), nil
	}
	return -1, nil
}

func (pr *PatientRepo) Report() (map[string]float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	sumStage := bson.D{{"$addFields",
		bson.D{{"total", bson.D{{"$add",
			bson.D{{"$sum", "$therapy.price"}},
		}},
		}},
	}}
	projectStage := bson.D{{"$project",
		bson.D{{"name", 1}, {"surname", 1}, {"total", 1}},
	}}

	cursor, err := patientsCollection.Aggregate(ctx, mongo.Pipeline{sumStage, projectStage})
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	report := make(map[string]float64)
	for _, result := range results {
		pr.logger.Println(result)
		report[result["_id"].(primitive.ObjectID).Hex()] = result["total"].(float64)
	}
	return report, nil
}

func (pr *PatientRepo) getCollection() *mongo.Collection {
	patientDatabase := pr.cli.Database("mongoDemo")
	patientsCollection := patientDatabase.Collection("patients")
	return patientsCollection
}

func (ur *UserRepo) getCollection() *mongo.Collection {
	dbName := "mongoDemo"
	usersCollectionName := "users"
	userDatabase := ur.cli.Database(dbName)

	test := os.Getenv("TEST")

	if test == "YES" {
		usersCollection := userDatabase.Collection(usersCollectionName + "_test")
		return usersCollection
	}
	usersCollection := userDatabase.Collection("users")
	return usersCollection
}
