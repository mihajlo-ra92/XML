package startup

import (
	"fmt"
	"log"

	"github.com/mihajlo-ra92/XML/accommodation_service/domain"
	"github.com/mihajlo-ra92/XML/accommodation_service/infrastructure/persistence"
	accommodation "github.com/mihajlo-ra92/XML/common/proto/accommodation_service"

	"github.com/mihajlo-ra92/XML/accommodation_service/application"
	"github.com/mihajlo-ra92/XML/accommodation_service/startup/config"

	"github.com/mihajlo-ra92/XML/accommodation_service/infrastructure/api"

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
	accommodationStore := server.initAccommodationStore(mongoClient)
	accommodationService := server.initAccommodationService(accommodationStore)
	accommodationHandler := server.initAccommodationHandler(accommodationService)
	server.startGrpcServer(accommodationHandler)

}


func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AccommodationDBHost, server.config.AccommodationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAccommodationStore(client *mongo.Client) domain.AccommodationStore {
	store := persistence.NewAccommodationMongoDBStore(client)
	store.DeleteAll()
	for _, Accommodation := range accommodations {
		err := store.Insert(Accommodation)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initAccommodationService(store domain.AccommodationStore) *application.AccommodationService {
	return application.NewAccommodationService(store)
}

// func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscribe

func (server *Server) initAccommodationHandler(service *application.AccommodationService) *api.AccommodationHandler{
	return api.NewAccommodationHandler(service)
}

func (server *Server) startGrpcServer(accommodationHandler *api.AccommodationHandler){
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accommodation.RegisterAccommodationServiceServer(grpcServer, accommodationHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}