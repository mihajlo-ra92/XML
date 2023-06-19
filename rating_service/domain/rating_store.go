package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingStore interface {
	Get(id primitive.ObjectID) (*Rating, error)
	GetAll() ([]*Rating, error)
	Insert(user *Rating) error
	DeleteAll()
	Delete(rating *Rating) error
	GetUserRatingByAccommodationId(accommodationId string, guestId string) (*Rating, error)
	GetUserRatingByHostId(hostId string, guestId string) (*Rating, error)
	GetAerageRatingByHostId(hostId string) (float32, error)
}
