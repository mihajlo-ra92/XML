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

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	bookingEndpoint := fmt.Sprintf("%s:%s", server.config.BookingHost, server.config.BookingPort)
	ratingEndpoint := fmt.Sprintf("%s:%s", server.config.RatingHost, server.config.RatingPort)
	notificationsEndpoint := fmt.Sprintf("%s:%s", server.config.NotificationsHost, server.config.NotificationsPort)
	authService := server.initAuthService(userEndpoint, accommodationEndpoint, bookingEndpoint, ratingEndpoint, notificationsEndpoint)
	authHandler := server.initAuthHandler(authService)
	server.startGrpcServer(authHandler)
}

func (server *Server) initAuthService(userClientAddress string, accommodationClientAddress string, bookingClientAddress string, ratingClientAddress string, notificationsClientAddress string) *application.AuthService {
	return application.NewAuthService(userClientAddress, accommodationClientAddress, bookingClientAddress, ratingClientAddress, notificationsClientAddress)
}

func (server *Server) initAuthHandler(service *application.AuthService) *api.AuthHandler {
	return api.NewAuthHandler(service)
}

// func (server *Server) initUserConn() error {
// 	conn, err := grpc.Dial()
// }

func (server *Server) startGrpcServer(authHandler *api.AuthHandler) {
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
