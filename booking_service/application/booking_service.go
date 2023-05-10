package application

import (
	"github.com/mihajlo-ra92/XML/booking_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingService struct {
	store domain.BookingStore
}

func NewBookingService(store domain.BookingStore) *BookingService {
	return &BookingService{
		store: store,
	}
}

func (service *BookingService) Get(id primitive.ObjectID) (*domain.Booking, error) {
	return service.store.Get(id)
}

func (service *BookingService) GetAll() ([]*domain.Booking, error) {
	return service.store.GetAll()
}

func (service *BookingService) Create(booking *domain.Booking) error {
	return service.store.Insert(booking)
}

func (service *BookingService) Reserve(booking *domain.Booking) error {

	return service.store.Insert(booking)
}
