package persistence

import (
	"context"
	"fmt"
	"log"

	accommodation "github.com/mihajlo-ra92/XML/common/proto/accommodation_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient(host, port string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	options := options.Client().ApplyURI(uri)
	return mongo.Connect(context.TODO(), options)
}

func NewAccommodationClient(address string) accommodation.AccommodationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return accommodation.NewAccommodationServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
