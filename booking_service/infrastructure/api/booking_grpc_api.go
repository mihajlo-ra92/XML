package api

import (
	"context"
	"fmt"
	"time"

	"github.com/mihajlo-ra92/XML/booking_service/application"
	"github.com/mihajlo-ra92/XML/booking_service/domain"
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

func (handler *BookingHandler) DeleteBooking(ctx context.Context, request *pb.DeleteBookingRequest) (*pb.DeleteBookingResponse, error) {
	fmt.Println("In DeleteBooking grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)
	err := handler.service.Delete(request.BookingId)
	if err != nil {
		return nil, err
	}
	response := pb.DeleteBookingResponse{DeletedBooking: &pb.Booking{}}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *BookingHandler) DeleteBookingsByGuestId(ctx context.Context, request *pb.DeleteBookingByGuestIdRequest) (*pb.DeleteBookingByGuestIdResponse, error) {
	fmt.Println("In DeleteBookingsByUserId grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)
	err := handler.service.DeleteByGuestId(request.UserId)
	if err != nil {
		return nil, err
	}
	response := pb.DeleteBookingByGuestIdResponse{Message: "Bookings deleted"}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *BookingHandler) DeleteBookingsByAccommodationId(ctx context.Context, request *pb.DeleteBookingByAccommodationIdRequest) (*pb.DeleteBookingByAccommodationIdResponse, error) {
	fmt.Println("In DeleteBookingsByAccommodationId grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)
	err := handler.service.DeleteByAccommodationId(request.AccommodationId)
	if err != nil {
		return nil, err
	}
	response := pb.DeleteBookingByAccommodationIdResponse{Message: "Bookings deleted"}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *BookingHandler) CreateBooking(ctx context.Context, request *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
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
	fmt.Println("In GetByAccomodationIdandDataRange grpc api")
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

func (handler *BookingHandler) GetBookingByAccommodationAndGuestId(ctx context.Context, request *pb.GetBookingByAccommodationAndGuestIdRequest) (*pb.GetBookingByAccommodationAndGuestIdResponse, error) {
	fmt.Println("In GetBookingByAccommodationAndGuestId grpc api")
	bookings, err := handler.service.GetByAccomodationAndGuestId(request.AccommodationId, request.GuestId)

	if err != nil {
		return nil, err
	}
	response := &pb.GetBookingByAccommodationAndGuestIdResponse{
		Bookings: []*pb.Booking{},
	}
	for _, booking := range bookings {
		current := mapBooking(booking)

		fmt.Println("Ispis bookinga: ", current)

		response.Bookings = append(response.Bookings, current)
	}
	return response, nil
}

func (handler *BookingHandler) ReservationCanceling(ctx context.Context, request *pb.ReservationCancelingRequest) (*pb.ReservationCancelingResponse, error) {
	fmt.Println("In ReservationCanceling grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)

	fmt.Print("booking after mapping: ")
	objectID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	resbook, err := handler.service.Get(objectID)
	if err != nil {
		return nil, err
	}

	res, err := handler.service.ReservationCanceling(resbook)
	if err != nil {
		return nil, err
	}
	response := pb.ReservationCancelingResponse{Booking: mapBooking(res)}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *BookingHandler) GetAllByUserAndType(ctx context.Context, request *pb.GetAllByUserRequest) (*pb.GetAllByUserResponse, error) {
	fmt.Println("In ReservationCanceling grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)

	bookings, err := handler.service.GetAllByUser(request.Id, domain.BookingType(request.BookingType))
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllByUserResponse{
		Bookings: []*pb.Booking{},
	}
	for _, booking := range bookings {
		current := mapBooking(booking)
		response.Bookings = append(response.Bookings, current)
	}
	return response, nil
}

func (handler *BookingHandler) GetByAccommodationId(ctx context.Context, request *pb.GetByAccommodationIdRequest) (*pb.GetByAccommodationIdResponse, error) {
	fmt.Println("In GetByAccommodationId grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)

	bookings, err := handler.service.GetByAccommodationId(request.AccommodationId)
	if err != nil {
		return nil, err
	}
	fmt.Print(bookings)
	response := &pb.GetByAccommodationIdResponse{Bookings: []*pb.Booking{}}
	for _, booking := range bookings {
		current := mapBooking(booking)
		response.Bookings = append(response.Bookings, current)
	}
	return response, nil
}

func (handler *BookingHandler) GetCancellationRateForHost(ctx context.Context, request *pb.GetCancellationRateForHostRequest) (*pb.GetCancellationRateForHostResponse, error) {
	fmt.Println("In GetCancellationRateForHost grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)

	percentage, err := handler.service.GetCancellationRateForHost(request.HostId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetCancellationRateForHostResponse{Percentage: percentage}

	return response, nil
}

func (handler *BookingHandler) GetNumberPastBookingsForHost(ctx context.Context, request *pb.GetNumberPastBookingsForHostRequest) (*pb.GetNumberPastBookingsForHostResponse, error) {
	fmt.Println("In GetNumberPastBookingsForHost grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)

	percentage, err := handler.service.GetNumberPastBookingsForHost(request.HostId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetNumberPastBookingsForHostResponse{Number: percentage}

	return response, nil
}
