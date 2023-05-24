package services

import (
	"log"

	accommodation "github.com/mihajlo-ra92/XML/common/proto/accommodation_service"
	booking "github.com/mihajlo-ra92/XML/common/proto/booking_service"
	rating "github.com/mihajlo-ra92/XML/common/proto/rating_service"
	user "github.com/mihajlo-ra92/XML/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(address string) user.UserServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return user.NewUserServiceClient(conn)
}

func NewAccommodationClient(address string) accommodation.AccommodationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return accommodation.NewAccommodationServiceClient(conn)
}

func NewBookingClient(address string) booking.BookingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Booking service: %v", err)
	}
	return booking.NewBookingServiceClient(conn)
}

func NewRatingClient(address string) rating.RatingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Rating service: %v", err)
	}
	return rating.NewRatingServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
