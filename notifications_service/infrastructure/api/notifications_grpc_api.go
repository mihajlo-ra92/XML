package api

import (
	"context"
	"fmt"

	pb "github.com/mihajlo-ra92/XML/common/proto/notifications_service"
	"github.com/mihajlo-ra92/XML/common/saga/messaging/nats"
	"github.com/mihajlo-ra92/XML/notifications_service/application"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationsHandler struct {
	pb.UnimplementedNotificationsServiceServer
	service *application.NotificationsService
}

func NewNotificationsHandler(service *application.NotificationsService) *NotificationsHandler {
	return &NotificationsHandler{
		service: service,
	}
}

func (handler *NotificationsHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	rating, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	ratingPb := mapNotification(rating)
	response := &pb.GetResponse{
		Rating: ratingPb,
	}
	return response, nil
}

/*
	func (handler *NotificationsHandler) GetConnection(ctx context.Context, request *pb.GetRequest) (*pb.Response, error) {
		userID := request.Id

		// Kreiranje WebSocket dialera
		dialer := websocket.Dialer{}

		// Postavljanje opcija za WebSocket konekciju
		// Ovdje možete postaviti bilo koje dodatne opcije, npr. autentifikacija
		// options := websocket.Dialer{}.SetHeader("Authorization", "Bearer <token>")

		// Establiciranje WebSocket konekcije
		conn, _, err := dialer.Dial("ws://localhost:8000/ws", nil)
		if err != nil {
			log.Println("WebSocket dial error:", err)
			return nil, err
		}

		// Dodavanje veze u mapu aktivnih veza korisnika
		Clients[userID] = conn

		// Ostatak vaše logike ovdje...

		return &pb.Response{
			Message: "Connection successful",
		}, nil

}
*/
func (handler *NotificationsHandler) GetByUser(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	rating, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	ratingPb := mapNotification(rating)
	response := &pb.GetResponse{
		Rating: ratingPb,
	}
	return response, nil
}

func (handler *NotificationsHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	fmt.Println("In GetAll grpc api")
	ratings, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Ratings: []*pb.Notification{},
	}
	for _, rating := range ratings {
		current := mapNotification(rating)
		response.Ratings = append(response.Ratings, current)
	}
	return response, nil
}

func (handler *NotificationsHandler) CreateNotification(ctx context.Context, request *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	fmt.Println("In CreateNotification grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)
	rating := mapNewNotification(request)
	fmt.Print("notification after mapping: ")
	fmt.Println(rating)
	err := handler.service.Create(rating)
	if err != nil {
		return nil, err
	}
	response := pb.CreateNotificationResponse{Notification: mapNotification(rating)}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *NotificationsHandler) DeleteNotification(ctx context.Context, request *pb.DeleteNotificationRequest) (*pb.DeleteNotificationResponse, error) {
	fmt.Println("In DeleteNotification grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)

	err := handler.service.Delete(request.NotificationId)
	if err != nil {
		return nil, err
	}

	response := pb.DeleteNotificationResponse{}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

/*
	func SendMessageToUser(userID string, message string) error {
		// Povežite se na NATS server
		nc, err := nats.Connect(nats.DefaultURL) /*NATS_HOST=nats
		NATS_PORT=4222
		NATS_USER=ruser
		NATS_PASS=T0pS3cr3t
		if err != nil {
			return err
		}
		defer nc.Close()

		// Generirajte jedinstvenu temu za poruku na temelju ID-a korisnika
		topic := fmt.Sprintf("user.%s", userID)

		// Pošaljite poruku na generiranu temu
		err = nc.Publish(topic, []byte(message))
		if err != nil {
			return err
		}

		// Provjerite jesu li poruke uspješno poslane
		nc.Flush()
		if err := nc.LastError(); err != nil {
			return err
		}

		return nil
	}
*/
func (handler *NotificationsHandler) SendMessage(ctx context.Context, request *pb.SendRequest) (*pb.SendResponse, error) {
	userID := request.Id
	message := request.Message
	nc, err := nats.NewPCublisher("nats", "4222", "ruser", "T0pS3cr3t", "app") //  .getConnection("nats", "4222", "ruser", "T0pS3cr3t") //nats.Connect(nats.DefaultURL)
	ntas.NewPCublisher()
	if err != nil {
		return nil, err
	}
	defer nc.Close()

	// Generirajte jedinstvenu temu za poruku na temelju ID-a korisnika
	topic := fmt.Sprintf("user.%s", userID)

	// Pošaljite poruku na generiranu temu
	err = nc.Publish(topic, []byte(message))
	if err != nil {
		return nil, err
	}

	// Provjerite jesu li poruke uspješno poslane
	nc.Flush()
	if err := nc.LastError(); err != nil {
		return nil, err
	}

	response := pb.SendResponse{Message: "Success"}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}
