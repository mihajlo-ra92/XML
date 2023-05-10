package application

import (
	"fmt"

	"github.com/mihajlo-ra92/XML/user_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(id primitive.ObjectID) (*domain.User, error){
	return service.store.Get(id)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) GetByLoginData(username string, password string) (*domain.User, error) {
	user, err := service.store.GetByLoginData(username, password)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, fmt.Errorf("incorrect password")
	}
	return user, nil
}

func (service *UserService) Create(user *domain.User) error{
	// user.UserType = domain.Guest
	checkUser, err := service.store.GetByUsername(user.Username)
	if err != nil && err.Error() != "mongo: no documents in result" {
	    return err
	}
	fmt.Print("Get user by username: ")
	fmt.Println(checkUser)
	if checkUser != nil {
		return fmt.Errorf("user with username %s already exists", user.Username)
	}
	
	checkUser, err = service.store.GetByEmail(user.Email)
	if err != nil && err.Error() != "mongo: no documents in result" {
	    return err
	}
	if checkUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	fmt.Print("Get user by email: ")
	fmt.Println(checkUser)
	//TODO: Optional 
	err = service.store.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) Update(user *domain.User) error{
	checkUser, err := service.store.GetByUsername(user.Username)
	if err != nil && err.Error() != "mongo: no documents in result" {
	    return err
	}
	fmt.Print("Get user by username: ")
	fmt.Println(checkUser)
	if checkUser != nil && checkUser.Id != user.Id {
		return fmt.Errorf("user with username %s already exists", user.Username)
	}
	
	checkUser, err = service.store.GetByEmail(user.Email)
	if err != nil && err.Error() != "mongo: no documents in result" {
	    return err
	}
	if checkUser != nil && checkUser.Id != user.Id  {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	fmt.Print("Get user by email: ")
	fmt.Println(checkUser)
	err = service.store.Delete(user)
	if err != nil {
		return err
	}
	err = service.store.Insert(user)
	if err != nil {
		return err
	}
	return nil
}