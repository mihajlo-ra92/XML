package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationsStore interface {
	Get(id primitive.ObjectID) (*Notification, error)
	GetByUser(userId string) (*Notification, error)
	GetAll() ([]*Notification, error)
	Insert(notification *Notification) error
	DeleteAll()
	Delete(notification *Notification) error
}
