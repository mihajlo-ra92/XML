package api

import (
	"context"
	"fmt"
	"log"

	"github.com/gorilla/websocket"

	"github.com/mihajlo-ra92/XML/notifications_service/application"

	pb "github.com/mihajlo-ra92/XML/common/proto/notifications_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Clients = make(map[string]*websocket.Conn)

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

func (handler *NotificationsHandler) SendMessage(ctx context.Context, request *pb.SendRequest) (*pb.SendResponse, error) {
	userID := request.Id
	message := request.Message

	conn, ok := Clients[userID]
	if !ok {
		log.Println("User is not connected:", userID)
		err1 := fmt.Errorf("User is not connected")
		return nil, err1
	}

	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("WebSocket write error:", err)
		err2 := fmt.Errorf("Failed to send message")
		return nil, err2
	}

	response := pb.SendResponse{Message: "Success"}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func sendNotification(userID string, message []byte) {
	conn, ok := Clients[userID]
	if !ok {
		log.Println("User is not connected:", userID)
		return
	}

	err := conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("WebSocket write error:", err)
	}
}
