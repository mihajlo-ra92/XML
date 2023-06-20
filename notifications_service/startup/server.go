package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/nats-io/nats.go"

	notification "github.com/mihajlo-ra92/XML/common/proto/notifications_service"
	"github.com/mihajlo-ra92/XML/notifications_service/application"
	"github.com/mihajlo-ra92/XML/notifications_service/domain"
	"github.com/mihajlo-ra92/XML/notifications_service/infrastructure/persistence"
	"github.com/mihajlo-ra92/XML/notifications_service/startup/config"
	"google.golang.org/grpc"

	"github.com/mihajlo-ra92/XML/notifications_service/infrastructure/api"

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
	notificationsStore := server.initNotificationsStore(mongoClient)
	notificationsService := server.initNotificationsService(notificationsStore)
	notificationsHandler := server.initNotificationsHandler(notificationsService)
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

func (server *Server) initRatingHandler(service *application.NotificationsService) *api.NotificationsHandler {
	return api.NewNotificationsHandler(service)
}

func (server *Server) startGrpcServer(ratingHandler *api.NotificationsHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	notification.RegisterNotificationsServiceServer(grpcServer, ratingHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	go StartMessageListener()
	fmt.Println("Messaging nats")
}

func StartMessageListener() {
	// Povežite se na NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Kreirajte kanal za primanje poruka
	msgChan := make(chan *nats.Msg)

	// Pretplatite se na teme na kojima želite primati poruke
	// Ovdje možete dodati više pretplata prema potrebama vašeg mikroservisa
	sub, err := nc.ChanSubscribe("user.*", msgChan)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	// Obrada primljenih poruka
	for msg := range msgChan {
		// Dobijte ID korisnika iz teme poruke
		topic := msg.Subject
		userID := topic[len("user."):]

		// Ovdje možete implementirati logiku obrade primljene poruke prema ID-u korisnika
		fmt.Printf("Primljena poruka za korisnika %s: %s\n", userID, string(msg.Data))
	}
}
