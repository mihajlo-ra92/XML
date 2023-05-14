package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationStore interface {
	Get(id primitive.ObjectID) (*Accommodation, error)
	GetByHostId(hostId string) ([]*Accommodation, error)
	GetAll() ([]*Accommodation, error)
	Insert(accommotaion *Accommodation) error
	DeleteAll()
	DeleteByHostId(hostId string) error
	Search(location string, numberGuests uint32) ([]*Accommodation, error)
	// Update()
}
