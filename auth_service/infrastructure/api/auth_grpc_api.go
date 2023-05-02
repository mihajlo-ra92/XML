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
	fmt.Print("Request.Login")
	fmt.Println(request.Login)
	return nil, nil
	// res := handler.service.Login(request.Login.Username, request.Login.Password)
}