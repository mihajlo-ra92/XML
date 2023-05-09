package api

import (
	"context"
	"fmt"

	"github.com/mihajlo-ra92/XML/booking_service/application"

	pb "github.com/mihajlo-ra92/XML/common/proto/booking_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingHandler struct {
	pb.UnimplementedBookingServiceServer
	service *application.BookingService
}

func NewBookingHandler(service *application.BookingService) *BookingHandler {
	return &BookingHandler{
		service: service,
	}
}

func (handler *BookingHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	booking, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	bookingPb := mapBooking(booking)
	response := &pb.GetResponse{
		Booking: bookingPb,
	}
	return response, nil
}

func (handler *BookingHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	fmt.Println("In GetAll grpc api")
	bookings, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Bookings: []*pb.Booking{},
	}
	for _, booking := range bookings {
		current := mapBooking(booking)
		response.Bookings = append(response.Bookings, current)
	}
	return response, nil
}

func (handler *BookingHandler) CreateBooking(ctx context.Context, request *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	//TODO: Implement
	fmt.Println("In CreateBooking grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)
	booking := mapNewBooking(request)
	fmt.Print("booking after mapping: ")
	fmt.Println(booking)
	err := handler.service.Create(booking)
	if err != nil {
		return nil, err
	}
	response := pb.CreateBookingResponse{Booking: mapBooking(booking)}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}