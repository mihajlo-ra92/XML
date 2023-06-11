package startup

import (
	"github.com/mihajlo-ra92/XML/booking_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var bookings = []*domain.Booking{
	{
		Id:              getObjectId("test1_id"),
		AccommodationId: "accommodation1Id",
		GuestId:         "guest1Id",
		Price:           10,
		PriceType:       domain.Regular,
		NumberOfGuests:  2,
		BookingType:     domain.Booked,
	},
	{
		Id:              getObjectId("test12_id"),
		AccommodationId: "accommodation1Id",
		GuestId:         "guest1Id",
		Price:           10,
		PriceType:       domain.Regular,
		NumberOfGuests:  2,
		BookingType:     domain.Reserved,
	},
	{
		Id:              getObjectId("test123_id"),
		AccommodationId: "accommodation1Id",
		GuestId:         "guest1Id",
		Price:           10,
		PriceType:       domain.Regular,
		NumberOfGuests:  2,
		BookingType:     domain.Reserved,
	},
	{
		Id:              getObjectId("test1234_id"),
		AccommodationId: "accommodation1Id",
		GuestId:         "guest1Id",
		Price:           10,
		PriceType:       domain.Regular,
		NumberOfGuests:  2,
		BookingType:     domain.Reserved,
	},
	{
		Id:              getObjectId("test12345_id"),
		AccommodationId: "accommodation1Id",
		GuestId:         "guest1Id",
		Price:           10,
		PriceType:       domain.Regular,
		NumberOfGuests:  2,
		BookingType:     domain.Reserved,
	},
	{
		Id:              getObjectId("test13456_id"),
		AccommodationId: "accommodation1Id",
		GuestId:         "guest1Id",
		Price:           10,
		PriceType:       domain.Regular,
		NumberOfGuests:  2,
		BookingType:     domain.Canceled,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
