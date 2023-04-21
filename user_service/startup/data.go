package startup

import (
	"github.com/mihajlo-ra92/XML/user_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
var users = []*domain.User{
	{
		Id:				getObjectId("test1_id"),
		UserType: 		domain.UserType,
		Username:		"test1",
		Password:		"123",
		Email:			"test1@gmail.com",
		FirstName:		"Test1F",
		LastName:		"Test1L",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID
}