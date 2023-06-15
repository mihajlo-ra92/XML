package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	Id                    primitive.ObjectID `bson:"_id"`
	UserId                string             `bson:"user_id"`
	HostRequest           string             `bson:"host_request"`
	HostCancelation       string             `bson:"host_cancelation"`
	HostRate              string             `bson:"host_rate"`
	HostAccommodationRate string             `bson:"host_accommodation_rate"`
	HostOutstanding       string             `bson:"host_outstanding"`
	GuestRequest          string             `bson:"guest_request"`
}
