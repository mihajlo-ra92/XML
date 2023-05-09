package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Accommodation struct {
	Id        primitive.ObjectID `bson:"_id"`
	HostId	  string		     `bson:"host_id"`
	Name      string             `bson:"name"`
	Location  string             `bson:"location"`
	Benefits  []string           `bson:"benefits"`
	Pictures  []string           `bson:"pictures"`
	MinGuests uint32             `bson:"min_guests"`
	MaxGuests uint32             `bson:"max_guests"`
}
