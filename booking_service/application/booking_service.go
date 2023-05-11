package application

import (
	"time"

	"fmt"

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

func TimeSpansOverlap(start1, end1, start2, end2 time.Time) bool {
	return start1.Before(end2) && end1.After(start2)
}

func (service *BookingService) Reserve(booking *domain.Booking) error {

	return service.store.Insert(booking)
}

func (service *BookingService) Deny(booking *domain.Booking) error {
	return service.store.Delete(booking)
}

func (service *BookingService) Book(booking *domain.Booking) error {
	if booking.StartDate.Before(time.Now()) {
		return fmt.Errorf("this date is before today's date")
	}
	if booking.EndDate.Before(booking.StartDate) {
		return fmt.Errorf("end date is before start date")
	}
	bookings, err := service.store.GetAll()
	if err != nil {
		return err
	}

	for _, oneBooking := range bookings {

		if oneBooking.AccommodationId == booking.AccommodationId {
			if TimeSpansOverlap(oneBooking.StartDate, oneBooking.EndDate, booking.StartDate, booking.EndDate) && oneBooking.BookingType == 2 {
				return fmt.Errorf("this date is already booked")
			}
		}
	}

	err = service.store.Delete(booking)
	if err != nil {
		return err
	}
	return service.store.Insert(booking)
}
