package application

import (
	"context"
	"fmt"

	"github.com/mihajlo-ra92/XML/accommodation_service/domain"
	"github.com/mihajlo-ra92/XML/accommodation_service/infrastructure/persistence"
	pb "github.com/mihajlo-ra92/XML/common/proto/accommodation_service"
	booking "github.com/mihajlo-ra92/XML/common/proto/booking_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	store                domain.AccommodationStore
	bookingClientAddress string
}

func NewAccommodationService(store domain.AccommodationStore, bookingClientAddress string) *AccommodationService {
	return &AccommodationService{
		store:                store,
		bookingClientAddress: bookingClientAddress,
	}
}

func (service *AccommodationService) Get(id primitive.ObjectID) (*domain.Accommodation, error) {
	return service.store.Get(id)
}

func (service *AccommodationService) GetAll() ([]*domain.Accommodation, error) {
	return service.store.GetAll()
}

func (service *AccommodationService) Create(accommodation *domain.Accommodation) error {
	return service.store.Insert(accommodation)
}

func (service *AccommodationService) Search(request *pb.SearchRequest) ([]*domain.Accommodation, error) {
	fmt.Println("In Search accommodation_service")
	bookingClient := persistence.NewBookingClient(service.bookingClientAddress)
	fmt.Println("booking_free_accomodation_search:")
	accommodations, err := service.store.Search(request.Location, request.Guest)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if (request.StartDate == nil) && (request.EndDate == nil) {
		return accommodations, nil
	}
	response := []*domain.Accommodation{}
	for _, accommodation := range accommodations {
		fmt.Println("accommodation id je  ", accommodation.Id.Hex())
		bookingResponse, err := bookingClient.GetByAccomodationIdandDataRange(context.TODO(), &booking.GetByAccomodationIdandDataRangeRequest{Id: accommodation.Id.Hex(), StartDate: request.StartDate, EndDate: request.EndDate})
		fmt.Println(bookingResponse)

		fmt.Println("udje ovde da se vidi booking")

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		var a uint32
		a = 0
		fmt.Println("da vidimo listu ; ", bookingResponse.Bookings)

		for _, booking := range bookingResponse.Bookings {
			fmt.Println("udje ovde da se vidi booking ; ", booking)

			a = a + booking.NumberOfGuests
			fmt.Println("sabiramo a")

		}

		fmt.Println("vrednost a : ", a)
		fmt.Println(accommodation.MaxGuests - request.Guest)

		if (accommodation.MaxGuests - request.Guest) >= a {
			response = append(response, accommodation)
		}

	}

	fmt.Println("Work here : ")
	fmt.Println("booking search for free accomodation Response: ")
	fmt.Println("create search accommodation response: ")
	return response, nil

}
