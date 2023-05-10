package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PriceType int8
type BookingType int8

const (
	Regular PriceType = iota
	PerGuest
)

const (
	CustomPrice BookingType = iota
	Reserved
	Booked
)

type Booking struct {
	Id              primitive.ObjectID `bson:"_id"`
	AccommodationId string             `bson:"accommodation_id"`
	GuestId         string             `bson:"guest_id"`
	Price           uint32             `bson:"price"`
	PriceType       PriceType          `bson:"price_type"`
	NumberOfGuests  uint32             `bson:"number_of_guests"`
	BookingType     BookingType        `bson:"booking_type"`
	StartDate       time.Time          `bson:"start_date"`
	EndDate         time.Time          `bson:"end_date"`
}
