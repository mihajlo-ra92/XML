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

type TicketRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

func (fr *TicketRepo) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbName := "mongoDemo"
	ticketsCollectionName := "tickets"
	ticketDatabase := fr.cli.Database(dbName)

	test := os.Getenv("TEST")

	if test == "YES" {
		ticketsCollection := ticketDatabase.Collection(ticketsCollectionName + "_test")
		ticketsCollection.Drop(ctx)
		return nil
	}
	return nil
}

func NewTicketRepo(ctx context.Context, logger *log.Logger) (*TicketRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &TicketRepo{
		cli:    client,
		logger: logger,
	}, nil
}

func (fr *TicketRepo) DisconnectTicketRepo(ctx context.Context) error {
	err := fr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ur *TicketRepo) Ping() {
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

func (tr *TicketRepo) getCollection() *mongo.Collection {
	dbName := "mongoDemo"
	ticketsCollectionName := "tickets"
	ticketDatabase := tr.cli.Database(dbName)

	test := os.Getenv("TEST")

	if test == "YES" {
		ticketsCollection := ticketDatabase.Collection(ticketsCollectionName + "_test")
		return ticketsCollection
	}
	ticketsCollection := ticketDatabase.Collection("tickets")
	return ticketsCollection
}

func (tr *TicketRepo) CheckSeats(ticket *Ticket, flight *Flight) bool {

	if flight == nil {
		tr.logger.Println("Non-existent flight")
		return false
	}

	if ticket.Capacity > flight.FreeSeats {
		tr.logger.Println("There are not enough seats")
		return false
	}

	if ticket.Capacity <= 0 {
		tr.logger.Println("There are not enough seats")
		return false
	}

	tr.logger.Printf("there are enough seats")
	return true
}

func (tr *TicketRepo) Insert(ticket *Ticket) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ticketsCollection := tr.getCollection()

	result, err := ticketsCollection.InsertOne(ctx, &ticket)
	if err != nil {
		tr.logger.Println(err)
		return err
	}
	tr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (tr *TicketRepo) GetAll() (Tickets, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticketsCollection := tr.getCollection()

	var tickets Tickets
	ticketsCursor, err := ticketsCollection.Find(ctx, bson.M{})
	if err != nil {
		tr.logger.Println(err)
		return nil, err
	}
	if err = ticketsCursor.All(ctx, tickets); err != nil {
		tr.logger.Println(err)
		return nil, err
	}
	return tickets, nil
}

func (tr *TicketRepo) GetById(id string) (*Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticketsCollection := tr.getCollection()

	var ticket Ticket
	objID, _ := primitive.ObjectIDFromHex(id)
	err := ticketsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&ticket)
	if err != nil {
		tr.logger.Println(err)
		return nil, err
	}
	return &ticket, nil
}

func (tr *TicketRepo) GetByIUserId(id string) (Tickets, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//objectId, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "userId", Value: id}}

	ticketsCollectoin := tr.getCollection()
	var tickets Tickets

	ticketsCursor, err := ticketsCollectoin.Find(ctx, filter)

	if err != nil {
		tr.logger.Println(err)
		return nil, err
	}
	if err = ticketsCursor.All(ctx, &tickets); err != nil {
		tr.logger.Println(err)
		return nil, err
	}
	return tickets, nil
}

func (tr *TicketRepo) Update(id string, ticket *Ticket) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ticketsCollection := tr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"date":       ticket.Date,
		"endPlace":   ticket.EndPlace,
		"startPlace": ticket.StartPlace,
		"capacity":   ticket.Capacity,
		"price":      ticket.Price,
		"userId":     ticket.UserID,
	}}
	result, err := ticketsCollection.UpdateOne(ctx, filter, update)
	tr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	tr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		tr.logger.Println(err)
		return err
	}
	return nil
}

func (tr *TicketRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ticketsCollection := tr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := ticketsCollection.DeleteOne(ctx, filter)
	if err != nil {
		tr.logger.Println(err)
		return err
	}
	tr.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}
