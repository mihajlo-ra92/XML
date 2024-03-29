package startup

import (
	"fmt"
	"log"

	user "github.com/mihajlo-ra92/XML/common/proto/user_service"
	"github.com/mihajlo-ra92/XML/user_service/domain"
	"github.com/mihajlo-ra92/XML/user_service/infrastructure/persistence"

	"github.com/mihajlo-ra92/XML/user_service/application"
	"github.com/mihajlo-ra92/XML/user_service/startup/config"

	"github.com/mihajlo-ra92/XML/user_service/infrastructure/api"

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
	userStore := server.initUserStore(mongoClient)
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	bookingEndpoint := fmt.Sprintf("%s:%s", server.config.BookingHost, server.config.BookingPort)
	userService := server.initUserService(userStore, accommodationEndpoint, bookingEndpoint)
	userHandler := server.initUserHandler(userService)
	server.startGrpcServer(userHandler)

}


func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	store.DeleteAll()
	for _, User := range users {
		err := store.Insert(User)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initUserService(store domain.UserStore, accommodationClientAddress string, bookingClientAddress string) *application.UserService {
	return application.NewUserService(store, accommodationClientAddress, bookingClientAddress)
}

// func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscribe

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler{
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler){
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}