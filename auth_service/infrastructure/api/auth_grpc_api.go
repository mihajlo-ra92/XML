package api

import (
	"context"
	"fmt"

	"github.com/mihajlo-ra92/XML/auth_service/application"
	pb "github.com/mihajlo-ra92/XML/common/proto/auth_service"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (handler *AuthHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error){
	fmt.Println("In Login grpc api")
	fmt.Print("Request.Login.Username: ")
	fmt.Println(request.Login.Username)
	fmt.Print("Request.Login.Password: ")
	fmt.Println(request.Login.Password)
	jwt, err := handler.service.Login(request.Login.Username, request.Login.Password)
	if err != nil {
		return nil, err
	}
	retVal := &pb.LoginResponse{Jwt: *jwt}
	return retVal, nil
	// res := handler.service.Login(request.Login.Username, request.Login.Password)
}