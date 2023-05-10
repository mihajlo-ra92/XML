package api

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mihajlo-ra92/XML/auth_service/application"
	"github.com/mihajlo-ra92/XML/auth_service/domain"
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

func (handler *AuthHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
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
}

func (handler *AuthHandler) AuthCreateAccommodation(ctx context.Context, request *pb.AuthCreateAccommodationRequest) (*pb.AuthCreateAccommodationResponse, error) {
	fmt.Println("In AuthCreateAccommodation")
	fmt.Print("request: ")
	fmt.Println(request)

	jwtData, err := checkJwt(request.Jwt)
	if err != nil {
		return nil, err
	}
	fmt.Print("jwtData: ")
	fmt.Println(jwtData)
	if jwtData.UserType != 1 {
		return nil, fmt.Errorf("user must be of host type")
	}
	accommodationResponse, err := handler.service.CreateAccommodation(jwtData, request)
	if err != nil {
		return nil, err
	}
	return accommodationResponse, nil
}

func (handler *AuthHandler) AuthUpdateUser(ctx context.Context, request *pb.AuthUpdateUserRequest) (*pb.AuthUpdateUserResponse, error){
	fmt.Println("In AuthUpdateUser")
	fmt.Print("request: ")
	fmt.Println(request)

	jwtData, err := checkJwt(request.Jwt)
	if err != nil {
		return nil, err
	}
	fmt.Print("jwtData: ")
	fmt.Println(jwtData)
	if jwtData.UserId != request.User.Id {
		return nil, fmt.Errorf("try to edit other user")
	}
	userResponse, err := handler.service.UpdateUser(request)
	if err != nil {
		return nil, err
	}
	return userResponse
}

func checkJwt(tokenString string) (*domain.JwtData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretKey := os.Getenv("SECRET_KEY")
		fmt.Println(secretKey)
		return []byte(secretKey), nil
	})
	fmt.Println("TOKEN: ")
	fmt.Println(token)
	var jwtData *domain.JwtData
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Valid jwt")
		fmt.Println(claims)
		jwtData = &domain.JwtData{
			Username: claims["username"].(string),
			UserId:   claims["userId"].(string),
			UserType: claims["userType"].(float64),
		}
	} else {
		fmt.Println(err)
		return nil, err
	}
	return jwtData, nil
}
