package api

import (
	"mihajlo-ra92/microservices_demo/user_service/domain"

	pb "mihajlo-ra92/microservices_demo/common/proto/user_service"
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