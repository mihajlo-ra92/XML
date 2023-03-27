package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	// NoSQL: module containing Mongo api client
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type FlightRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

func (fr *FlightRepo) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbName := "mongoDemo"
	flightsCollectionName := "flights"
	flightDatabase := fr.cli.Database(dbName)

	test := os.Getenv("TEST")

	if test == "YES" {
		flightsCollection := flightDatabase.Collection(flightsCollectionName + "_test")
		flightsCollection.Drop(ctx)
		return nil
	}
	return nil
}

func NewFlightRepo(ctx context.Context, logger *log.Logger) (*FlightRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &FlightRepo{
		cli:    client,
		logger: logger,
	}, nil
}

func (fr *FlightRepo) DisconnectFlightRepo(ctx context.Context) error {
	err := fr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ur *FlightRepo) Ping() {
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

func (fr *FlightRepo) getCollection() *mongo.Collection {
	dbName := "mongoDemo"
	flightsCollectionName := "flights"
	flightDatabase := fr.cli.Database(dbName)

	test := os.Getenv("TEST")

	if test == "YES" {
		flightsCollection := flightDatabase.Collection(flightsCollectionName + "_test")
		return flightsCollection
	}
	flightsCollection := flightDatabase.Collection("flights")
	return flightsCollection
}

func (fr *FlightRepo) Insert(flight *Flight) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flightsCollection := fr.getCollection()

	result, err := flightsCollection.InsertOne(ctx, &flight)
	if err != nil {
		fr.logger.Println(err)
		return err
	}
	fr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (fr *FlightRepo) GetAll() (Flights, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := fr.getCollection()

	var flights Flights
	flightsCursor, err := flightsCollection.Find(ctx, bson.M{})
	if err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &flights); err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	return flights, nil
}

func (fr *FlightRepo) GetFlightById(id string) (Flights, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}

	flightsCollectoin := fr.getCollection()
	var flights Flights

	flightsCursor, err := flightsCollectoin.Find(ctx, filter)
	// usersCursor, err := usersCollectoin.Find(ctx, bson.M{"username": username})

	if err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &flights); err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	return flights, nil
}

func (fr *FlightRepo) Update(id string, flight *Flight) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := fr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"date":       flight.Date,
		"endPlace":   flight.EndPlace,
		"startPlace": flight.StartPlace,
		"capacity":   flight.Capacity,
		"price":      flight.Price,
		"freeSeats":  flight.FreeSeats,
	}}
	result, err := usersCollection.UpdateOne(ctx, filter, update)
	fr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	fr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		fr.logger.Println(err)
		return err
	}
	return nil
}

func (fr *FlightRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flightsCollection := fr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := flightsCollection.DeleteOne(ctx, filter)
	if err != nil {
		fr.logger.Println(err)
		return err
	}
	fr.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}
