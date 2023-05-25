package application

import (
	"context"
	"fmt"

	"github.com/mihajlo-ra92/XML/rating_service/domain"
	"github.com/mihajlo-ra92/XML/rating_service/infrastructure/persistence"

	booking "github.com/mihajlo-ra92/XML/common/proto/booking_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingService struct {
	store                      domain.RatingStore
	accommodationClientAddress string
	bookingClientAddress       string
	userClientAddress          string
}

func NewRatingService(store domain.RatingStore, accommodationClientAddress string, bookingClientAddress string, userClientAddress string) *RatingService {
	return &RatingService{
		store:                      store,
		accommodationClientAddress: accommodationClientAddress,
		bookingClientAddress:       bookingClientAddress,
		userClientAddress:          userClientAddress,
	}
}

func (service *RatingService) Get(id primitive.ObjectID) (*domain.Rating, error) {
	return service.store.Get(id)
}

func (service *RatingService) GetAll() ([]*domain.Rating, error) {
	return service.store.GetAll()
}

func (service *RatingService) Create(rating *domain.Rating) error {

	fmt.Print("Rating for creating: ")
	fmt.Println(rating)
	if rating.Rate < 1 || rating.Rate > 5 {
		return fmt.Errorf("Rates must be between 1 and 5")
	}
	bookingClient := persistence.NewBookingClient(service.bookingClientAddress)
	bookingRequest := booking.GetBookingByAccommodationAndGuestIdRequest{AccommodationId: rating.AccommodationId, GuestId: rating.GuestId}
	bookingResponse, err := bookingClient.GetBookingByAccommodationAndGuestId(context.TODO(), &bookingRequest)

	fmt.Print("bookingResponse: ")
	fmt.Println(bookingResponse)

	if bookingResponse == nil {
		return fmt.Errorf("The user hasn't been in this accommodation")
	}
	//TODO: Optional
	err = service.store.Insert(rating)
	if err != nil {
		return err
	}
	return nil
}
