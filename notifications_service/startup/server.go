package startup

import (
	"fmt"
	"log"

	//notifications "github.com/mihajlo-ra92/XML/common/proto/notifications_service"
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
	notificationsStore := server.initRatingStore(mongoClient)
	notificationsService := server.initRatingService(notificationsStore)
	notificationsHandler := server.initRatingHandler(notificationsService)
	server.startGrpcServer(notificationsHandler)

}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.NotificationsDBHost, server.config.NotificationsDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initNotificationsStore(client *mongo.Client) domain.NotificationsStore {
	store := persistence.NewNotificationsMongoDBStore(client)
	store.DeleteAll()
	for _, Notifications := range notifications {
		err := store.Insert(Notifications)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initNotificationsService(store domain.NotificationsStore) *application.NotificationsService {
	return application.NewNotificationsService(store)
}

func (server *Server) initNotificationsHandler(service *application.NotificationsService) *api.NotificationsHandler {
	return api.NewNotificationsHandler(service)
}

func (server *Server) startGrpcServer(notificationsHandler *api.NotificationsHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	rating.RegisterRatingServiceServer(grpcServer, notificationsHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
