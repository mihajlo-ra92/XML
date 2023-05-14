package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingStore interface {
	Get(id primitive.ObjectID) (*Booking, error)
	GetAll() ([]*Booking, error)
	GetAllByUser(guestId string, bookingType BookingType) ([]*Booking, error)
	Insert(booking *Booking) error
	DeleteAll()
	Update(booking *Booking) (*Booking, error)
	Delete(booking *Booking) error
	GetByAccomodationIdandDataRange(accommodationId string, startDate time.Time, endDate time.Time) ([]*Booking, error)
	// Update()
}
