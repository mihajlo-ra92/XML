package startup

import (
	"github.com/mihajlo-ra92/XML/rating_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ratings = []*domain.Rating{
	{
		Id:              getObjectId("rate1"),
		HostId:          "host1Id",
		AccommodationId: "accommodation1Id",
		GuestId:         "guest1Id",
		Rate:            4,
	},
	{
		Id:              getObjectId("rate12"),
		HostId:          "host1Id",
		AccommodationId: "accommodation1Id",
		GuestId:         "guest1Id",
		Rate:            1,
	},
	{
		Id:              getObjectId("rate123"),
		HostId:          "host1Id",
		AccommodationId: "accommodation1Id",
		GuestId:         "guest1Id",
		Rate:            3,
	},
	{
		Id:              getObjectId("rate1234"),
		HostId:          "host12Id",
		AccommodationId: "accommodation1Id",
		GuestId:         "guest1Id",
		Rate:            3,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
