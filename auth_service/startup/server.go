package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/mihajlo-ra92/XML/auth_service/application"
	"github.com/mihajlo-ra92/XML/auth_service/infrastructure/api"
	"github.com/mihajlo-ra92/XML/auth_service/startup/config"
	auth "github.com/mihajlo-ra92/XML/common/proto/auth_service"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server{
	return &Server{
		config: config,
	}
}

func (server *Server) Start(){
	//mongoClient
	//authStore
	authService := server.initAuthService()
	authHandler := server.initAuthHandler(authService)
	server.startGrpcServer()
}

func (server *Server) initAuthService() *application.AuthService {
	return application.NewAuthService()
}

func (server *Server) initAuthHandler(service *application.AuthService) *api.AuthHandler{
	return api.NewAuthHandler(service)
}

func (server *Server) startGrpcServer(authHandler *api.AuthHandler){
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, authHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}