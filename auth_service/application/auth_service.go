package application

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mihajlo-ra92/XML/auth_service/infrastructure/services"
	user "github.com/mihajlo-ra92/XML/common/proto/user_service"
)

type AuthService struct {
	userClientAddress          string
	accommodationClientAddress string
}

func NewAuthService(userClientAddress string, accommodationClientAddress string) *AuthService {
	return &AuthService{
		userClientAddress:          userClientAddress,
		accommodationClientAddress: accommodationClientAddress,
	}
}

func (service *AuthService) Login(username string, password string) (*string, error) {
	fmt.Println("In auth_service, login")
	userClient := services.NewUserClient(service.userClientAddress)
	dataToSend := user.Login{Username: username, Password: password}
	fmt.Print("dataToSend: ")
	fmt.Println(dataToSend)
	userResp, err := userClient.GetByLoginData(context.TODO(), &user.GetByLoginDataRequest{Login: &dataToSend})
	if err != nil {
		return nil, err
	}
	fmt.Print("Read user: ")
	fmt.Print(userResp.User)

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userResp.User.Username,
		"userType": userResp.User.UserType,
		"userId":   userResp.User.Id,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	secretKey := os.Getenv("SECRET_KEY")
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

// func (service *AuthService) ReadJwt(jwt string) (*domain.User, error){
// 	return nil, nil
// }
