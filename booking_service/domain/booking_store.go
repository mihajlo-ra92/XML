package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingStore interface {
	Get(id primitive.ObjectID) (*Booking, error)
	GetAll() ([]*Booking, error)
	Insert(booking *Booking) error
	DeleteAll()
	// Update()
}
