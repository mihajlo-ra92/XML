package api

import (
	pb "github.com/mihajlo-ra92/XML/common/proto/user_service"
	"github.com/mihajlo-ra92/XML/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapNewUser(userPb *pb.NewUser) *domain.User {
	user := &domain.User{

		Id:          primitive.NewObjectID(),
		UserType:    domain.UserType(userPb.UserType),
		Username:    userPb.Username,
		Password:    userPb.Password,
		Email:       userPb.Email,
		FirstName:   userPb.FirstName,
		LastName:    userPb.LastName,
		Address:     userPb.Address,
		Outstanding: "NO",
	}
	return user
}

func mapUpdatedUser(userPb *pb.User) *domain.User {
	UserId, _ := primitive.ObjectIDFromHex(userPb.Id)
	user := &domain.User{
		Id:          UserId,
		UserType:    domain.UserType(userPb.UserType),
		Username:    userPb.Username,
		Password:    userPb.Password,
		Email:       userPb.Email,
		FirstName:   userPb.FirstName,
		LastName:    userPb.LastName,
		Address:     userPb.Address,
		Outstanding: userPb.Outstanding,
	}
	return user
}

func mapUser(user *domain.User) *pb.User {
	var userType pb.User_UserType
	switch user.UserType {
	case domain.Admin:
		userType = pb.User_Admin
	case domain.Guest:
		userType = pb.User_Guest
	case domain.Host:
		userType = pb.User_Host
	}
	userPb := &pb.User{
		Id:          user.Id.Hex(),
		UserType:    userType,
		Username:    user.Username,
		Password:    user.Password,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Address:     user.Address,
		Outstanding: user.Outstanding,
	}
	return userPb
}
