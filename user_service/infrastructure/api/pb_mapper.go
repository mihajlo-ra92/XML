package api

import (
	pb "github.com/mihajlo-ra92/XML/common/proto/user_service"
	"github.com/mihajlo-ra92/XML/user_service/domain"
	// pb "github.com/mihajlo-ra92/XML/common/proto/user_service"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:		user.Id.Hex(),
		UserType: 	user.UserType,
		Username:	user.Username,
		Password:	user.Password,
		Email: 		user.Email,
		FirstName:	user.FirstName,
		LastName:	user.LastName,
	}
	return userPb
}