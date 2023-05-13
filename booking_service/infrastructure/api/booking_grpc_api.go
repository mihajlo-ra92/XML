package api

import (
	"context"
	"fmt"
	"time"

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

func (handler *BookingHandler) GuestReserveAccommodation(ctx context.Context, request *pb.GuestReserveAccommodationRequest) (*pb.GuestReserveAccommodationResponse, error) {
	fmt.Println("In GuestReserveAccommodation grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)

	reservation := mapNewReservation(request)
	fmt.Print("booking after mapping: ")
	fmt.Println(reservation)
	err := handler.service.Reserve(reservation)
	if err != nil {
		return nil, err
	}

	response := pb.GuestReserveAccommodationResponse{Booking: mapBooking(reservation)}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *BookingHandler) BookingAccept(ctx context.Context, request *pb.BookingAcceptRequest) (*pb.BookingAcceptResponse, error) {
	fmt.Println("In BookingAccept grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)

	if request.Booking.BookingType != 1 {
		return nil, fmt.Errorf("this accommodation isn't reserved")
	}

	reservation := mapAcceptedBooking(request)
	fmt.Print("booking after mapping: ")
	fmt.Println(reservation)
	err := handler.service.Book(reservation)
	if err != nil {
		return nil, err
	}

	response := pb.BookingAcceptResponse{Booking: mapBooking(reservation)}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *BookingHandler) BookingDeny(ctx context.Context, request *pb.BookingDenyRequest) (*pb.BookingDenyResponse, error) {
	fmt.Println("In BookingDeny grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)

	if request.Booking.BookingType != 1 {
		return nil, fmt.Errorf("this accommodation isn't reserved")
	}

	reservation := mapDeniedBooking(request)
	fmt.Print("booking after mapping: ")
	fmt.Println(reservation)
	err := handler.service.Deny(reservation)
	if err != nil {
		return nil, err
	}

	response := pb.BookingDenyResponse{}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *BookingHandler) GetByAccomodationIdandDataRange(ctx context.Context, request *pb.GetByAccomodationIdandDataRangeRequest) (*pb.GetByAccomodationIdandDataRangeResponse, error) {
	fmt.Println("In GetAll grpc api")
	bookings, err := handler.service.GetByAccomodationIdandDataRange(request.Id, time.Unix(request.StartDate.Seconds, int64(request.StartDate.Nanos)).UTC(), time.Unix(request.EndDate.Seconds, int64(request.EndDate.Nanos)).UTC())

	if err != nil {
		return nil, err
	}
	response := &pb.GetByAccomodationIdandDataRangeResponse{
		Bookings: []*pb.Booking{},
	}
	for _, booking := range bookings {
		current := mapBooking(booking)

		fmt.Println("Ispis bookinga: ", current)

		response.Bookings = append(response.Bookings, current)
	}
	return response, nil
}
