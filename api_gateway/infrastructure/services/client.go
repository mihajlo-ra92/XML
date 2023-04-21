package services

import (
	"log"

	catalogue "github.com/mihajlo-ra92/XML/common/proto/catalogue_service"
	ordering "github.com/mihajlo-ra92/XML/common/proto/ordering_service"
	shipping "github.com/mihajlo-ra92/XML/common/proto/shipping_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCatalogueClient(address string) catalogue.CatalogueServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return catalogue.NewCatalogueServiceClient(conn)
}

func NewOrderingClient(address string) ordering.OrderingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Ordering service: %v", err)
	}
	return ordering.NewOrderingServiceClient(conn)
}

func NewShippingClient(address string) shipping.ShippingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Shipping service: %v", err)
	}
	return shipping.NewShippingServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
