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

	fmt.Println("Work here : ")
	bookingResponse, err := bookingClient.GetAll(context.TODO(), &booking.GetAllRequest{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("booking search for free accomodation Response: ")
	fmt.Println(bookingResponse)
	fmt.Println("create search accommodation response: ")
	return service.store.Search(request.Location, request.Guest)
}
