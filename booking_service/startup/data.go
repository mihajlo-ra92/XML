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
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
