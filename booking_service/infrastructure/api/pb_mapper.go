package api

import (
	"github.com/mihajlo-ra92/XML/booking_service/domain"
	pb "github.com/mihajlo-ra92/XML/common/proto/booking_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapBooking(booking *domain.Booking) *pb.Booking {
	var priceType pb.Booking_PriceType
	switch booking.PriceType {
	case domain.Regular:
		priceType = pb.Booking_Regular
	case domain.PerGuest:
		priceType = pb.Booking_PerGuest
	}

	var bookingType pb.Booking_BookingType
	switch booking.BookingType {
	case domain.Booked:
		bookingType = pb.Booking_Booked
	case domain.CustomPrice:
		bookingType = pb.Booking_CustomPrice
	case domain.Reserved:
		bookingType = pb.Booking_Reserved
	}
	bookingPb := &pb.Booking{
		Id:              booking.Id.Hex(),
		AccommodationId: booking.AccommodationId,
		GuestId:         booking.GuestId,
		Price:           booking.Price,
		PriceType:       priceType,
		BookingType:     bookingType,
		NumberOfGuests:  booking.NumberOfGuests,
		StartDate:       timestamppb.New(booking.StartDate),
		EndDate:         timestamppb.New(booking.EndDate),
	}
	return bookingPb
}

func mapNewBooking(request *pb.CreateBookingRequest) *domain.Booking {
	booking := &domain.Booking{

		Id:              primitive.NewObjectID(),
		AccommodationId: request.Booking.AccommodationId,
		GuestId:         request.Booking.GuestId,
		Price:           request.Booking.Price,
		PriceType:       domain.PriceType(request.Booking.PriceType),
		NumberOfGuests:  request.Booking.NumberOfGuests,
		BookingType:     domain.BookingType(request.Booking.BookingType),
	}
	return booking
}

func mapNewReservation(request *pb.GuestReserveAccommodationRequest) *domain.Booking {
	booking := &domain.Booking{

		Id:              primitive.NewObjectID(),
		AccommodationId: request.Booking.AccommodationId,
		GuestId:         request.Booking.GuestId,
		Price:           request.Booking.Price,
		PriceType:       domain.PriceType(request.Booking.PriceType),
		NumberOfGuests:  request.Booking.NumberOfGuests,
		BookingType:     1,
		StartDate:       request.Booking.StartDate.AsTime(),
		EndDate:         request.Booking.EndDate.AsTime(),
	}
	return booking
}

func mapAcceptedBooking(request *pb.BookingAcceptRequest) *domain.Booking {
	BookingId, _ := primitive.ObjectIDFromHex(request.Booking.Id)

	booking := &domain.Booking{

		Id:              BookingId,
		AccommodationId: request.Booking.AccommodationId,
		GuestId:         request.Booking.GuestId,
		Price:           request.Booking.Price,
		PriceType:       domain.PriceType(request.Booking.PriceType),
		NumberOfGuests:  request.Booking.NumberOfGuests,
		BookingType:     2,
		StartDate:       request.Booking.StartDate.AsTime(),
		EndDate:         request.Booking.EndDate.AsTime(),
	}
	return booking
}
