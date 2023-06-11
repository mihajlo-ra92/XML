package application

import (
	"context"
	"fmt"
	"time"

	"github.com/mihajlo-ra92/XML/booking_service/domain"
	"github.com/mihajlo-ra92/XML/booking_service/infrastructure/persistence"
	accommodation "github.com/mihajlo-ra92/XML/common/proto/accommodation_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingService struct {
	store                      domain.BookingStore
	accommodationClientAddress string
}

func NewBookingService(store domain.BookingStore, accommodationClientAddress string) *BookingService {
	return &BookingService{
		store:                      store,
		accommodationClientAddress: accommodationClientAddress,
	}
}

func (service *BookingService) Get(id primitive.ObjectID) (*domain.Booking, error) {
	return service.store.Get(id)
}

func (service *BookingService) GetAll() ([]*domain.Booking, error) {
	return service.store.GetAll()
}

func (service *BookingService) Delete(bookingId string) error {
	id, _ := primitive.ObjectIDFromHex(bookingId)
	booking := domain.Booking{Id: id}
	return service.store.Delete(&booking)
}

func (service *BookingService) DeleteByGuestId(guestId string) error {
	return service.store.DeleteByGuestId(guestId)
}

func (service *BookingService) DeleteByAccommodationId(accommodationId string) error {
	return service.store.DeleteByAccommodationId(accommodationId)
}

func (service *BookingService) Create(booking *domain.Booking) error {
	//OPTIMISATION: Implement special endpoint for defining custom price
	if booking.BookingType == domain.CustomPrice {
		//OPTIMISATION: implement get by accommodationId
		bookings, err := service.store.GetAll()
		if err != nil {
			return err
		}
		for _, bookingIt := range bookings {
			if bookingIt.AccommodationId == booking.AccommodationId {
				if TimeSpansOverlap(bookingIt.StartDate, bookingIt.EndDate, booking.StartDate, booking.EndDate) {
					if bookingIt.BookingType == domain.Booked || bookingIt.BookingType == domain.Reserved {
						return fmt.Errorf("given date range is taken")
					}
					if bookingIt.BookingType == domain.CustomPrice {
						service.store.Delete(bookingIt)
					}
				}
			}
		}
	}
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

func (service *BookingService) GetByAccomodationIdandDataRange(accommodationId string, startDate time.Time, endDate time.Time) ([]*domain.Booking, error) {
	return service.store.GetByAccomodationIdandDataRange(accommodationId, startDate, endDate)
}

func (service *BookingService) GetByAccomodationAndGuestId(accommodationId string, guestId string) ([]*domain.Booking, error) {
	return service.store.GetByAccomodationAndGuestId(accommodationId, guestId)
}

func (service *BookingService) ReservationCanceling(booking *domain.Booking) (*domain.Booking, error) {
	if time.Now().Before(booking.StartDate) {
		booking.BookingType = domain.Canceled
		fmt.Println("Radil ovo ", booking.BookingType)
	} else {
		return nil, fmt.Errorf("reservation is started")
	}
	return service.store.Update(booking)
}

func (service *BookingService) GetAllByUser(guestId string, bookingType domain.BookingType) ([]*domain.Booking, error) {
	return service.store.GetAllByUser(guestId, bookingType)
}

func (service *BookingService) GetByAccommodationId(accommodationId string) ([]*domain.Booking, error) {
	return service.store.GetByAccommodationId(accommodationId)
}

func (service *BookingService) GetCancellationRateForHost(hostId string) (uint32, error) {

	accommodationClient := persistence.NewAccommodationClient(service.accommodationClientAddress)
	accommodations, err := accommodationClient.GetByHostId(context.TODO(), &accommodation.GetByHostIdRequest{HostId: hostId})
	fmt.Println("Dosao je do linije 144")

	if err != nil {
		return 0, err
	}
	numberOfCancelation := 0
	sumBookingsByHost := 0
	fmt.Println("Dosao je do linije 151")

	for _, accommodation := range accommodations.Acccommodations {
		fmt.Println(accommodation)
		bookings, err := service.store.GetCancellationBookingsByAccommodation("accommodation1Id")
		fmt.Println(bookings)

		if err != nil {
			return 0, err
		}
		sumBookings, err := service.store.GetByAccommodationId("accommodation1Id")

		sumBookingsByHost = sumBookingsByHost + len(sumBookings)
		numberOfCancelation = numberOfCancelation + len(bookings)

	}
	if sumBookingsByHost == 0 {
		sumBookingsByHost = 1
	}

	fmt.Println(sumBookingsByHost, "Ovo je suma ")
	fmt.Println(numberOfCancelation, "Ovo je broj otkazanih")

	percentage := (float64(numberOfCancelation) / float64(sumBookingsByHost)) * 100
	fmt.Println((numberOfCancelation/sumBookingsByHost)*100, " Ovo je pravi procenat")

	fmt.Println(uint32(percentage))

	return uint32(percentage), nil
}

//func (service *RatingService) GetAerageRatingByHostId(hostId string) (float32, error) {
