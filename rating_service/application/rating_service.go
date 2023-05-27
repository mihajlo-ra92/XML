package application

import (
	"context"
	"fmt"

	"github.com/mihajlo-ra92/XML/rating_service/domain"
	"github.com/mihajlo-ra92/XML/rating_service/infrastructure/persistence"

	accommodation "github.com/mihajlo-ra92/XML/common/proto/accommodation_service"
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

func (service *RatingService) Delete(ratingId string) error {
	id, _ := primitive.ObjectIDFromHex(ratingId)
	rating := domain.Rating{Id: id}
	return service.store.Delete(&rating)
}

func (service *RatingService) GetUserRatingByAccommodationId(accommodationId string, guestId string) (*domain.Rating, error) {
	return service.store.GetUserRatingByAccommodationId(accommodationId, guestId)
}

func (service *RatingService) GetUserRatingByHostId(hostId string, guestId string) (*domain.Rating, error) {
	return service.store.GetUserRatingByHostId(hostId, guestId)
}

func (service *RatingService) Create(rating *domain.Rating) error {

	fmt.Print("Rating for creating: ")
	fmt.Println(rating)
	fmt.Println(rating.HostId)
	if rating.Rate < 1 || rating.Rate > 5 {
		return fmt.Errorf("Rates must be between 1 and 5")
	}
	bookingClient := persistence.NewBookingClient(service.bookingClientAddress)
	bookingRequest := booking.GetBookingByAccommodationAndGuestIdRequest{AccommodationId: rating.AccommodationId, GuestId: rating.GuestId}
	bookingResponse, err := bookingClient.GetBookingByAccommodationAndGuestId(context.TODO(), &bookingRequest)

	accommodationClient := persistence.NewAccommodationClient(service.accommodationClientAddress)
	accommodationRequest := accommodation.GetByHostIdRequest{HostId: rating.HostId}
	accommodationResponse, err := accommodationClient.GetByHostId(context.TODO(), &accommodationRequest)

	ratingForUpdateHost, err := service.store.GetUserRatingByHostId(rating.HostId, rating.GuestId)
	ratingForUpdateAccommodation, err := service.store.GetUserRatingByAccommodationId(rating.AccommodationId, rating.GuestId)

	if rating.AccommodationId == "" {
		if ratingForUpdateHost == nil {
			fmt.Print(rating)

			if accommodationResponse.Acccommodations != nil {

				for _, accommodation := range accommodationResponse.Acccommodations {
					bookingRequestForAccommodation := booking.GetBookingByAccommodationAndGuestIdRequest{AccommodationId: accommodation.Id, GuestId: rating.GuestId}
					bookingResponseForAccommodation, err := bookingClient.GetBookingByAccommodationAndGuestId(context.TODO(), &bookingRequestForAccommodation)
					fmt.Print("bookingResponseForAccommodation: ")
					fmt.Println(bookingResponseForAccommodation)

					if err != nil {
						return err
					}

					if bookingResponseForAccommodation.Bookings == nil {
						return fmt.Errorf("The guest hasn't been in any of this host's accommodations")
					}
					fmt.Print("bookingResponseForAccommodation.Bookings: ")
					fmt.Println(bookingResponseForAccommodation.Bookings)
				}
			} else {
				fmt.Println("This host doesn't have any accommodations")
			}
		} else {
			service.store.Delete(ratingForUpdateHost)
			err = service.store.Insert(rating)

			return fmt.Errorf("You have already rated this host, we will update your rate with this rating")

		}
	}

	if rating.HostId == "" {
		if ratingForUpdateAccommodation == nil {
			fmt.Print(rating)

			if bookingResponse.Bookings == nil {
				return fmt.Errorf("The guest hasn't been in this accommodation")
			}
		} else {
			fmt.Println("this is rating for update accommodation:")
			fmt.Println(rating)
			service.store.Delete(ratingForUpdateAccommodation)
			err = service.store.Insert(rating)

			return fmt.Errorf("You have already rated this accommodation, we will update your rate with this rating")
		}

	}

	fmt.Print("bookingResponse: ")
	fmt.Println(bookingResponse)
	fmt.Print("accommodationResponse: ")
	fmt.Println(accommodationResponse)

	//TODO: Optional
	err = service.store.Insert(rating)
	if err != nil {
		return err
	}
	return nil
}
