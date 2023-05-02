package application

import (
	"context"
	"fmt"

	"github.com/mihajlo-ra92/XML/auth_service/infrastructure/services"
	user "github.com/mihajlo-ra92/XML/common/proto/user_service"
)

type AuthService struct {
	userClientAddress string
}

func NewAuthService(userClientAddress string) *AuthService{
	return &AuthService{
		userClientAddress: userClientAddress,
	}
}

func (service *AuthService) Login(username string, password string) (*string, error){
	fmt.Println("In auth_service, login")
	userClient := services.NewUserClient(service.userClientAddress)
	dataToSend := user.Login{Username: username, Password: password}
	fmt.Print("dataToSend: ")
	fmt.Println(dataToSend)
	userResp, err :=userClient.GetByLoginData(context.TODO(), &user.GetByLoginDataRequest{Login: &dataToSend})
	if err != nil {
		return nil, err
	}
	fmt.Print("Read user: ")
	fmt.Print(userResp.User)
	
	return &userResp.User.Username, nil
}
