package application

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mihajlo-ra92/XML/auth_service/domain"
	"github.com/mihajlo-ra92/XML/auth_service/infrastructure/services"
	accommodation "github.com/mihajlo-ra92/XML/common/proto/accommodation_service"
	pb "github.com/mihajlo-ra92/XML/common/proto/auth_service"
	booking "github.com/mihajlo-ra92/XML/common/proto/booking_service"
	user "github.com/mihajlo-ra92/XML/common/proto/user_service"
)

type AuthService struct {
	userClientAddress          string
	accommodationClientAddress string
	bookingClientAddress       string
}

func NewAuthService(userClientAddress string, accommodationClientAddress string, bookingClientAddress string) *AuthService {
	return &AuthService{
		userClientAddress:          userClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		bookingClientAddress:       bookingClientAddress,
	}
}

func (service *AuthService) Login(username string, password string) (*string, error) {
	fmt.Println("In auth_service, login")
	userClient := services.NewUserClient(service.userClientAddress)
	fmt.Println(userClient)
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

func (service *AuthService) CreateAccommodation(jwtData *domain.JwtData, request *pb.AuthCreateAccommodationRequest) (*pb.AuthCreateAccommodationResponse, error) {
	accommodationClient := services.NewAccommodationClient(service.accommodationClientAddress)
	accommodationUser := accommodation.AccommodationUser{Id: jwtData.UserId, UserType: accommodation.AccommodationUser_UserType(jwtData.UserType), Username: jwtData.Username}
	accommodationRequest := accommodation.CreateAccommodationRequest{User: &accommodationUser, Name: request.Name, Location: request.Location, Benefits: request.Benefits, Pictures: request.Pictures, MinGuests: request.MinGuests, MaxGuests: request.MaxGuests}
	fmt.Print("accommodationRequest: ")
	fmt.Println(accommodationRequest)
	accommodationResponse, err := accommodationClient.CreateAccommodation(context.TODO(), &accommodationRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("accommodationResponse: ")
	fmt.Println(accommodationResponse)
	authCreateAccommodationResponse := pb.AuthCreateAccommodationResponse{Accomodation: &pb.Accommodation{Id: accommodationResponse.Accommodation.Id, HostId: accommodationResponse.Accommodation.HostId, Name: accommodationResponse.Accommodation.Name, Location: accommodationResponse.Accommodation.Location, Benefits: accommodationResponse.Accommodation.Benefits, Pictures: accommodationResponse.Accommodation.Pictures, MinGuests: accommodationResponse.Accommodation.MinGuests, MaxGuests: accommodationResponse.Accommodation.MaxGuests}}
	fmt.Print("authCreateAccommodationResponse: ")
	fmt.Println(authCreateAccommodationResponse)
	return &authCreateAccommodationResponse, nil
}

func (service *AuthService) UpdateUser(request *pb.AuthUpdateUserRequest) (*pb.AuthUpdateUserResponse, error) {
	userClient := services.NewUserClient(service.userClientAddress)
	userA := user.User{
		Id:        request.User.Id,
		UserType:  user.User_UserType(request.User.UserType),
		Username:  request.User.Username,
		Password:  request.User.Password,
		Email:     request.User.Email,
		FirstName: request.User.FirstName,
		LastName:  request.User.LastName,
		Address:   request.User.Address,
	}

	userUpdateRequest := user.UpdateUserRequest{User: &userA}
	fmt.Print("userUpdateRequest: ")
	fmt.Println(userUpdateRequest)
	userUpdateResponse, err := userClient.UpdateUser(context.TODO(), &userUpdateRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("userUpdateResponse: ")
	fmt.Println(userUpdateResponse)
	userB := pb.AuthUser{
		Id:        userUpdateResponse.User.Id,
		UserType:  pb.AuthUser_UserType(userUpdateResponse.User.UserType),
		Username:  userUpdateResponse.User.Username,
		Password:  userUpdateResponse.User.Password,
		Email:     userUpdateResponse.User.Email,
		FirstName: userUpdateResponse.User.FirstName,
		LastName:  userUpdateResponse.User.LastName,
		Address:   userUpdateResponse.User.Address,
	}
	authUpdateUserResponse := pb.AuthUpdateUserResponse{User: &userB}
	fmt.Print("authUpdateUserResponse: ")
	fmt.Println(authUpdateUserResponse)
	return &authUpdateUserResponse, nil
}

func (service *AuthService) DeleteUser(request *pb.AuthDeleteUserRequest) (*pb.AuthDeleteUserResponse, error) {
	userClient := services.NewUserClient(service.userClientAddress)
	userDeleteRequest := user.DeleteUserRequest{Id: request.Id}
	fmt.Print("userDeleteRequest: ")
	fmt.Println(userDeleteRequest)
	userDeleteResponse, err := userClient.DeleteUser(context.TODO(), &userDeleteRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("userDeleteResponse: ")
	fmt.Println(userDeleteResponse)
	userB := pb.AuthUser{
		Id:        userDeleteResponse.User.Id,
		UserType:  pb.AuthUser_UserType(userDeleteResponse.User.UserType),
		Username:  userDeleteResponse.User.Username,
		Password:  userDeleteResponse.User.Password,
		Email:     userDeleteResponse.User.Email,
		FirstName: userDeleteResponse.User.FirstName,
		LastName:  userDeleteResponse.User.LastName,
		Address:   userDeleteResponse.User.Address,
	}

	authDeleteUserResposne := pb.AuthDeleteUserResponse{User: &userB}
	fmt.Print("authDeleteUserResposne: ")
	fmt.Println(authDeleteUserResposne)
	return &authDeleteUserResposne, nil
}

func (service *AuthService) GuestReserveAccommodation(jwtData *domain.JwtData, request *pb.AuthGuestReserveAccommodationRequest) (*pb.AuthGuestReserveAccommodationResponse, error) {
	fmt.Println("In reserve accommodation")
	bookingClient := services.NewBookingClient(service.bookingClientAddress)
	bookingStruct := booking.Booking{AccommodationId: request.AccommodationId, GuestId: jwtData.UserId, Price: request.Price, PriceType: booking.Booking_PriceType(request.PriceType), NumberOfGuests: request.NumberOfGuests, BookingType: booking.Booking_BookingType(request.BookingType), StartDate: request.StartDate, EndDate: request.EndDate}
	bookingRequest := booking.GuestReserveAccommodationRequest{Booking: &bookingStruct}
	fmt.Print("bookingRequest: ")
	fmt.Println(bookingRequest)
	bookingResponse, err := bookingClient.GuestReserveAccommodation(context.TODO(), &bookingRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("bookingResponse: ")
	fmt.Println(bookingResponse)
	authReserveAccommodationResponse := pb.AuthGuestReserveAccommodationResponse{Booking: &pb.Booking{Id: bookingResponse.Booking.Id, AccommodationId: bookingResponse.Booking.AccommodationId, GuestId: bookingResponse.Booking.GuestId, Price: bookingResponse.Booking.Price, PriceType: pb.Booking_PriceType(bookingResponse.Booking.PriceType), NumberOfGuests: bookingResponse.Booking.NumberOfGuests, BookingType: pb.Booking_BookingType(bookingResponse.Booking.BookingType), StartDate: bookingResponse.Booking.StartDate, EndDate: bookingResponse.Booking.EndDate}}
	fmt.Print("authCreateBookingResponse: ")
	fmt.Println(authReserveAccommodationResponse)
	return &authReserveAccommodationResponse, nil
}

func (service *AuthService) BookingAccept(jwtData *domain.JwtData, request *pb.AuthBookingAcceptRequest) (*pb.AuthBookingAcceptResponse, error) {
	fmt.Println("In booking accept")
	bookingClient := services.NewBookingClient(service.bookingClientAddress)
	bookginGetRequest := booking.GetRequest{Id: request.BookingId}
	reservedBooking, err := bookingClient.Get(context.TODO(), &bookginGetRequest)
	if err != nil {
		return nil, err
	}
	bookingRequest := booking.BookingAcceptRequest{Booking: reservedBooking.Booking}
	fmt.Print("bookingRequest: ")
	fmt.Println(bookingRequest)
	bookingResponse, err := bookingClient.BookingAccept(context.TODO(), &bookingRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("bookingResponse: ")
	fmt.Println(bookingResponse)
	authBookingAcceptResponse := pb.AuthBookingAcceptResponse{Booking: &pb.Booking{Id: bookingResponse.Booking.Id, AccommodationId: bookingResponse.Booking.AccommodationId, GuestId: bookingResponse.Booking.GuestId, Price: bookingResponse.Booking.Price, PriceType: pb.Booking_PriceType(bookingResponse.Booking.PriceType), NumberOfGuests: bookingResponse.Booking.NumberOfGuests, BookingType: pb.Booking_BookingType(bookingResponse.Booking.BookingType), StartDate: bookingResponse.Booking.StartDate, EndDate: bookingResponse.Booking.EndDate}}
	fmt.Print("authBookingAcceptResponse: ")
	fmt.Println(authBookingAcceptResponse)
	return &authBookingAcceptResponse, nil
}

func (service *AuthService) BookingDeny(jwtData *domain.JwtData, request *pb.AuthBookingDenyRequest) (*pb.AuthBookingDenyResponse, error) {
	fmt.Println("In booking deny")
	bookingClient := services.NewBookingClient(service.bookingClientAddress)
	bookginGetRequest := booking.GetRequest{Id: request.BookingId}
	reservedBooking, err := bookingClient.Get(context.TODO(), &bookginGetRequest)
	if err != nil {
		return nil, err
	}
	bookingRequest := booking.BookingDenyRequest{Booking: reservedBooking.Booking}
	fmt.Print("bookingRequest: ")
	fmt.Println(bookingRequest)
	bookingResponse, err := bookingClient.BookingDeny(context.TODO(), &bookingRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("bookingResponse: ")
	fmt.Println(bookingResponse)
	authBookingDenyResponse := pb.AuthBookingDenyResponse{}
	fmt.Print("authBookingDenyResponse: ")
	fmt.Println(authBookingDenyResponse)
	return &authBookingDenyResponse, nil
}

func (service *AuthService) CancelingReservation(jwtData *domain.JwtData, request *pb.AuthReservationCancelingRequest) (*pb.AuthReservationCancelingResponse, error) {
	fmt.Println("In reserve accommodation")
	bookingRequest := booking.ReservationCancelingRequest{Id: request.Id}
	fmt.Print("cancelingRequest: ")
	fmt.Println(bookingRequest)
	fmt.Println("In booking deny")
	bookingClient := services.NewBookingClient(service.bookingClientAddress)

	bookingResponse, err := bookingClient.ReservationCanceling(context.TODO(), &bookingRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("bookingResponse: ")
	fmt.Println(bookingResponse)
	authCanceling := pb.AuthReservationCancelingResponse{Booking: &pb.Booking{Id: bookingResponse.Booking.Id, AccommodationId: bookingResponse.Booking.AccommodationId, GuestId: bookingResponse.Booking.GuestId, Price: bookingResponse.Booking.Price, PriceType: pb.Booking_PriceType(bookingResponse.Booking.PriceType), NumberOfGuests: bookingResponse.Booking.NumberOfGuests, BookingType: pb.Booking_BookingType(bookingResponse.Booking.BookingType), StartDate: bookingResponse.Booking.StartDate, EndDate: bookingResponse.Booking.EndDate}}
	fmt.Print("authCreateBookingResponse: ")
	fmt.Println(authCanceling)
	return &authCanceling, nil
}

func (service *AuthService) DefineCustomPrice(jwtData *domain.JwtData, request *pb.AuthDefineCustomPriceRequest) (*pb.AuthDefineCustomPriceResponse, error) {
	jwtUser := accommodation.AccommodationUser{Id: jwtData.UserId, UserType: accommodation.AccommodationUser_UserType(jwtData.UserType), Username: jwtData.Username}
	customPriceRequest := accommodation.DefineCustomPriceRequest{User: &jwtUser, AccommodationId: request.AccommodationId, StartDate: request.StartDate, EndDate: request.EndDate, Price: request.Price, PriceType: accommodation.DefineCustomPriceRequest_PriceType(request.PriceType)}
	fmt.Print("customPriceRequest: ")
	fmt.Println(customPriceRequest)
	fmt.Println("In booking deny")
	accommodationClient := services.NewAccommodationClient(service.accommodationClientAddress)

	accommodationResponse, err := accommodationClient.DefineCustomPrice(context.TODO(), &customPriceRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("accommodationResponse: ")
	fmt.Println(accommodationResponse)
	authDefineCustomPriceResponse := pb.AuthDefineCustomPriceResponse{Accommodation: &pb.Accommodation{Id: accommodationResponse.Accommodation.Id, HostId: accommodationResponse.Accommodation.HostId, Name: accommodationResponse.Accommodation.Name, Location: accommodationResponse.Accommodation.Location, Benefits: accommodationResponse.Accommodation.Benefits, Pictures: accommodationResponse.Accommodation.Pictures, MinGuests: accommodationResponse.Accommodation.MinGuests, MaxGuests: accommodationResponse.Accommodation.MaxGuests }}
	fmt.Print("authDefineCustomPriceResponse: ")
	fmt.Println(authDefineCustomPriceResponse)
	return &authDefineCustomPriceResponse, nil
}