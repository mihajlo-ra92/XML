package startup

import (
	"github.com/mihajlo-ra92/XML/rating_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var notifications = []*domain.Notification{
	{
		Id:                    getObjectId("rate1"),
		UserId:                "host1Id",
		HostRequest:           "NO",
		HostCancelation:       "NO",
		HostRate:              "NO",
		HostAccommodationRate: "NO",
		HostOutstanding:       "NO",
		GuestRequest:          "NO",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
