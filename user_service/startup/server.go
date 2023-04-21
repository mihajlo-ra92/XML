package startup

import (
	"fmt"
	"log"

	"github.com/mihajlo-ra92/XML/user_service/domain"
	"github.com/mihajlo-ra92/XML/user_service/persistence"

	"github.com/mihajlo-ra92/XML/user_service/application"
	"github.com/mihajlo-ra92/XML/user_service/config"

	"github.com/mihajlo-ra92/XML/user_service/api"

	user "github.com/mihajlo-ra92/XML/common/proto/user_service"

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

const (
	QueueGroup = "user_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)
	userService := server.initUserService(userStore)
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

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscribe

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
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}