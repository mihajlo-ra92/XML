package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserType int8

const (
	Guest UserType = iota
	Host
	Admin
)

func (userType UserType) String() string {
	switch userType {
	case Guest:
		return "Guest"
	case Host:
		return "Host"
	case Admin:
		return "Admin"
	}
	return "Unknown"
}

type User struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserType    UserType           `bson:"user_type"`
	Username    string             `bson:"username"`
	Password    string             `bson:"password"`
	Email       string             `bson:"email"`
	FirstName   string             `bson:"first_name"`
	LastName    string             `bson:"last_name"`
	Address     string             `bson:"address"`
	Outstanding string             `bson:"outstanding"`
}