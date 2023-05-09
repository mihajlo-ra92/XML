package startup

import (
	"fmt"
	"log"

	"github.com/mihajlo-ra92/XML/booking_service/domain"
	"github.com/mihajlo-ra92/XML/booking_service/infrastructure/persistence"
	booking "github.com/mihajlo-ra92/XML/common/proto/booking_service"

	"github.com/mihajlo-ra92/XML/booking_service/application"
	"github.com/mihajlo-ra92/XML/booking_service/startup/config"

	"github.com/mihajlo-ra92/XML/booking_service/infrastructure/api"

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
	bookingStore := server.initBookingStore(mongoClient)
	bookingService := server.initBookingService(bookingStore)
	bookingHandler := server.initBookingHandler(bookingService)
	server.startGrpcServer(bookingHandler)

}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.BookingDBHost, server.config.BookingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initBookingStore(client *mongo.Client) domain.BookingStore {
	store := persistence.NewBookingMongoDBStore(client)
	store.DeleteAll()
	for _, Booking := range bookings {
		err := store.Insert(Booking)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initBookingService(store domain.BookingStore) *application.BookingService {
	return application.NewBookingService(store)
}

// func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscribe

func (server *Server) initBookingHandler(service *application.BookingService) *api.BookingHandler {
	return api.NewBookingHandler(service)
}

func (server *Server) startGrpcServer(bookingHandler *api.BookingHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	booking.RegisterBookingServiceServer(grpcServer, bookingHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
