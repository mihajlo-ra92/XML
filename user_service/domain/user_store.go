package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Get(id primitive.ObjectID) (*User, error)
	GetAll() ([]*User, error)
	Insert(user *User) error
	DeleteAll()
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByLoginData(username string, password string) (*User, error)
	Update(user *User) error
	Delete(user *User) error
}