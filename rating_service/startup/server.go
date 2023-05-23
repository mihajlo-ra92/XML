package startup

import (
	"fmt"
	"log"

	rating "github.com/mihajlo-ra92/XML/common/proto/rating_service"
	"github.com/mihajlo-ra92/XML/rating_service/domain"
	"github.com/mihajlo-ra92/XML/rating_service/infrastructure/persistence"

	"github.com/mihajlo-ra92/XML/rating_service/application"
	"github.com/mihajlo-ra92/XML/rating_service/startup/config"

	"github.com/mihajlo-ra92/XML/rating_service/infrastructure/api"

	"net"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	ratingStore := server.initRatingStore(mongoClient)
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	bookingEndpoint := fmt.Sprintf("%s:%s", server.config.BookingHost, server.config.BookingPort)
	userEndpoint := fmt.Sprintf("%s:%s", server.config.BookingHost, server.config.BookingPort)
	ratingService := server.initRatingService(ratingStore, accommodationEndpoint, bookingEndpoint, userEndpoint)
	ratingHandler := server.initRatingHandler(ratingService)
	server.startGrpcServer(ratingHandler)

}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.RatingDBHost, server.config.RatingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initRatingStore(client *mongo.Client) domain.RatingStore {
	store := persistence.NewRatingMongoDBStore(client)
	store.DeleteAll()
	for _, Rating := range ratings {
		err := store.Insert(Rating)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initRatingService(store domain.RatingStore, accommodationClientAddress string, bookingClientAddress string, userClientAddress string) *application.RatingService {
	return application.NewRatingService(store, accommodationClientAddress, bookingClientAddress, userClientAddress)
}

// func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscribe

func (server *Server) initRatingHandler(service *application.RatingService) *api.RatingHandler {
	return api.NewRatingHandler(service)
}

func (server *Server) startGrpcServer(ratingHandler *api.RatingHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	rating.RegisterRatingServiceServer(grpcServer, ratingHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
