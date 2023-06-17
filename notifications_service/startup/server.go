package startup

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	notification "github.com/mihajlo-ra92/XML/common/proto/notifications_service"
	"github.com/mihajlo-ra92/XML/notifications_service/domain"
	"github.com/mihajlo-ra92/XML/notifications_service/infrastructure/persistence"

	"github.com/mihajlo-ra92/XML/notifications_service/application"
	"github.com/mihajlo-ra92/XML/notifications_service/startup/config"

	"github.com/mihajlo-ra92/XML/notifications_service/infrastructure/api"

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

func (server *Server) startGrpcServer(ratingHandler *api.NotificationsHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	notification.RegisterNotificationsServiceServer(grpcServer, ratingHandler)
	fmt.Println("Serving...")
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve gRPC: %s", err)
		}
	}()

	// Definisanje WebSocket endpointa
	http.HandleFunc("/ws", handleWebSocket)

	fmt.Println("Serving WebSocket...")

	// Pokretanje HTTP servera za WebSocket na istom portu
	if err := http.ListenAndServe(fmt.Sprintf(":%s", "50800"), nil); err != nil {
		log.Fatalf("failed to serve WebSocket: %s", err)
	}
}

/*
func (server *Server) startServer(ratingHandler *api.NotificationsHandler) {




	// Pokretanje gRPC servera u pozadini
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve gRPC: %s", err)
		}
	}()

	// Definisanje WebSocket endpointa
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Logika za WebSocket endpoint
	})

	fmt.Println("Serving WebSocket...")

	// Pokretanje HTTP servera za WebSocket na istom portu
	if err := http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), nil); err != nil {
		log.Fatalf("failed to serve WebSocket: %s", err)
	}
}
*/

func handleWebSocket(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("id")

	// Upgradirajte HTTP konekciju na WebSocket
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	api.Clients[userID] = conn

	go readWebSocketMessages(userID, conn)
}

func readWebSocketMessages(userID string, conn *websocket.Conn) {
	defer func() {
		delete(api.Clients, userID)
		conn.Close()
	}()

	for {

		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		log.Printf("Received message from user %s: %s\n", userID, string(message))
	}
}
