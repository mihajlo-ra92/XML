package api

import (
	pb "github.com/mihajlo-ra92/XML/common/proto/user_service"
	"github.com/mihajlo-ra92/XML/user_service/domain"
)

func mapUser(user *domain.User) *pb.User {
	var userType pb.User_UserType
	switch user.UserType{
	case domain.Admin:
		userType = pb.User_Admin
	case domain.Guest:
		userType = pb.User_Guest
	case domain.Host:
		userType = pb.User_Host
	}
	userPb := &pb.User{
		Id:		user.Id.Hex(),
		UserType: 	userType,
		Username:	user.Username,
		Password:	user.Password,
		Email: 		user.Email,
		FirstName:	user.FirstName,
		LastName:	user.LastName,
		Address:	user.Address,
	}
	return userPb
}