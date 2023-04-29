package startup

import (
	"github.com/mihajlo-ra92/XML/accommodation_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var accommodations = []*domain.Accommodation{
	{
		Id:				getObjectId("test1_id"),

	Name:      "Name1",
	Location: 	"Location1",
	Benefits:  []string{"Wifi", "Parking"},
	Pictures:  []string{"Wifi", "Parking"},
	MinGuests: 1,
	MaxGuests: 10,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}