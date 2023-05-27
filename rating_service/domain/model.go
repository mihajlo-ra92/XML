package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rating struct {
	Id              primitive.ObjectID `bson:"_id"`
	HostId          string             `bson:"host_id"`
	AccommodationId string             `bson:"accommodation_id"`
	GuestId         string             `bson:"guest_id"`
	Rate            uint32             `bson:"rate"`
}
