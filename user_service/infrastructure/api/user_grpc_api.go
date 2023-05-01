package api

import (
	"context"
	"fmt"

	"github.com/mihajlo-ra92/XML/user_service/application"

	pb "github.com/mihajlo-ra92/XML/common/proto/user_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	userPb := mapUser(user)
	response := &pb.GetResponse{
		User: userPb,
	}
	return response, nil
}

func (handler *UserHandler) GetByLoginData(ctx context.Context, request *pb.GetByLoginDataRequest) (*pb.GetByLoginDataResponse, error) {
	fmt.Println("In GetByLoginData grpc api")
	fmt.Println(request.Login)
	fmt.Print("request.Username: ")
	fmt.Println(request.Login.Username)
	fmt.Print("request.Password: ")
	fmt.Println(request.Login.Password)
	user, err := handler.service.GetByLoginData(request.Login.Username, request.Login.Password)
	if err != nil {
		return nil, err
	}
	userPb := mapUser(user)
	response := &pb.GetByLoginDataResponse{
		User: userPb,
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	fmt.Println("In GetAll grpc api")
	users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}
func (handler *UserHandler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Println("In CreateUser grpc api")
	fmt.Print("Request.User: ")
	fmt.Println(request.User)
	user := mapNewUser(request.User)
	fmt.Print("user after mapping: ")
	fmt.Println(user)
	err := handler.service.Create(user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		User: mapUser(user),
	}, nil
}
