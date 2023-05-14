package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type PriceType int8

const (
	Regular PriceType = iota
	PerGuest
)

func (priceType PriceType) String() string {
	switch priceType {
	case Regular:
		return "Regular"
	case PerGuest:
		return "PerGuest"
	}
	return "Unknown"
}

type Accommodation struct {
	Id        primitive.ObjectID `bson:"_id"`
	HostId    string             `bson:"host_id"`
	Name      string             `bson:"name"`
	Location  string             `bson:"location"`
	Benefits  []string           `bson:"benefits"`
	Pictures  []string           `bson:"pictures"`
	MinGuests uint32             `bson:"min_guests"`
	MaxGuests uint32             `bson:"max_guests"`
	Price     uint32             `bson:"price"`
	PriceType PriceType		     `bson:"price_type"`
}
