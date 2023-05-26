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
}
