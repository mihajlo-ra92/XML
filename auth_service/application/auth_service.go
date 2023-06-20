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
	noti "github.com/mihajlo-ra92/XML/common/proto/notifications_service"
	rating "github.com/mihajlo-ra92/XML/common/proto/rating_service"
	user "github.com/mihajlo-ra92/XML/common/proto/user_service"
)

type AuthService struct {
	userClientAddress          string
	accommodationClientAddress string
	bookingClientAddress       string
	ratingClientAddress        string
	notificationsClientAddress string
}

func NewAuthService(userClientAddress string, accommodationClientAddress string, bookingClientAddress string, ratingClientAddress string, notificationsClientAddress string) *AuthService {
	return &AuthService{
		userClientAddress:          userClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		bookingClientAddress:       bookingClientAddress,
		ratingClientAddress:        ratingClientAddress,
		notificationsClientAddress: notificationsClientAddress,
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
	accommodationRequest := accommodation.CreateAccommodationRequest{User: &accommodationUser, Name: request.Name, Location: request.Location, Benefits: request.Benefits, Pictures: request.Pictures, MinGuests: request.MinGuests, MaxGuests: request.MaxGuests, Price: request.Price, PriceType: accommodation.CreateAccommodationRequest_PriceType(request.PriceType)}
	fmt.Print("accommodationRequest: ")
	fmt.Println(accommodationRequest)
	accommodationResponse, err := accommodationClient.CreateAccommodation(context.TODO(), &accommodationRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("accommodationResponse: ")
	fmt.Println(accommodationResponse)
	authCreateAccommodationResponse := pb.AuthCreateAccommodationResponse{Accomodation: &pb.Accommodation{Id: accommodationResponse.Accommodation.Id, HostId: accommodationResponse.Accommodation.HostId, Name: accommodationResponse.Accommodation.Name, Location: accommodationResponse.Accommodation.Location, Benefits: accommodationResponse.Accommodation.Benefits, Pictures: accommodationResponse.Accommodation.Pictures, MinGuests: accommodationResponse.Accommodation.MinGuests, MaxGuests: accommodationResponse.Accommodation.MaxGuests, Price: accommodationRequest.Price, PriceType: pb.Accommodation_PriceType(accommodationResponse.Accommodation.PriceType)}}
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

func (service *AuthService) AuthCreateRating(jwtData *domain.JwtData, request *pb.AuthCreateRatingRequest) (*pb.AuthCreateRatingResponse, error) {
	fmt.Println("In create rating")
	ratingClient := services.NewRatingClient(service.ratingClientAddress)
	ratingStruct := rating.NewRating{HostId: request.HostId, AccommodationId: request.AccommodationId, GuestId: jwtData.UserId, Rate: request.Rate}
	ratingRequest := rating.CreateRatingRequest{Rating: &ratingStruct}
	guestId := jwtData.UserId
	fmt.Print(guestId)
	fmt.Print("ratingRequest: ")
	fmt.Println(ratingRequest)
	ratingResponse, err := ratingClient.CreateRating(context.TODO(), &ratingRequest)
	if err != nil {
		return nil, err
	}

	if err == nil {
		getAverageRatingByHostIdRequest := rating.GetAverageRatingByHostIdRequest{Id: request.HostId}

		ratingResponse1, err := ratingClient.GetAverageRatingByHostId(context.TODO(), &getAverageRatingByHostIdRequest)
		if err == nil {

			userClient := services.NewUserClient(service.userClientAddress)
			userGetRequest := user.GetRequest{Id: request.HostId}
			host, err := userClient.Get(context.TODO(), &userGetRequest)
			if err == nil {

				if ratingResponse1.AverageRating > 4.7 {
					if host.User.Outstanding != "YES" {
						host.User.Outstanding = "YES"
						userClient.UpdateUser(context.TODO(), &user.UpdateUserRequest{User: host.User})
					}
				} else {
					if host.User.Outstanding != "NO" {
						host.User.Outstanding = "NO"
						userClient.UpdateUser(context.TODO(), &user.UpdateUserRequest{User: host.User})
					}
				}
			}
		}
	}

	fmt.Print("ratingResponse: ")
	fmt.Println(ratingResponse)
	authCreateRatingResponse := pb.AuthCreateRatingResponse{Rating: &pb.Rating{Id: ratingResponse.Rating.Id, HostId: ratingResponse.Rating.HostId, AccommodationId: ratingResponse.Rating.AccommodationId, GuestId: guestId, Rate: ratingResponse.Rating.Rate}}
	fmt.Print("authCreateRatingResponse: ")
	fmt.Println(authCreateRatingResponse)
	return &authCreateRatingResponse, nil
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

	if err == nil {
		cancelingRequest := booking.GetCancellationRateForHostRequest{HostId: jwtData.UserId}
		numberRequest := booking.GetNumberPastBookingsForHostRequest{HostId: jwtData.UserId}

		numberResponse, err := bookingClient.GetNumberPastBookingsForHost(context.TODO(), &numberRequest)
		ratingResponse, err1 := bookingClient.GetCancellationRateForHost(context.TODO(), &cancelingRequest)
		if err == nil && err1 == nil {

			userClient := services.NewUserClient(service.userClientAddress)
			userGetRequest := user.GetRequest{Id: jwtData.UserId}
			host, err := userClient.Get(context.TODO(), &userGetRequest)
			if err == nil {

				if ratingResponse.Percentage < 5 && numberResponse.Number >= 5 {
					if host.User.Outstanding != "YES" {
						host.User.Outstanding = "YES"
						userClient.UpdateUser(context.TODO(), &user.UpdateUserRequest{User: host.User})
					}
				} else {
					if host.User.Outstanding != "NO" {
						host.User.Outstanding = "NO"
						userClient.UpdateUser(context.TODO(), &user.UpdateUserRequest{User: host.User})
					}
				}
			}
		}
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

func (service *AuthService) DeleteRating(jwtData *domain.JwtData, request *pb.AuthDeleteRatingRequest) (*pb.AuthDeleteRatingResponse, error) {
	ratingClient := services.NewRatingClient(service.ratingClientAddress)
	ratingDeleteRequest := rating.DeleteRatingRequest{Jwt: request.Jwt, RatingId: request.RatingId}
	fmt.Print("ratingDeleteRequest: ")
	fmt.Println(ratingDeleteRequest)
	reqe := noti.SendRequest{Id: "", Message: "Nova poruka za klijenta da je neko obrisao ocenu"}
	notClient := services.NewNotificationsClient(service.notificationsClientAddress)
	responsMessage, err := notClient.SendMessage(context.TODO(), &reqe)
	if err != nil {
		return nil, err
	}

	if responsMessage != nil {
	}

	ratingDeleteResponse, err := ratingClient.DeleteRating(context.TODO(), &ratingDeleteRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("ratingDeleteResponse: ")
	fmt.Println(ratingDeleteResponse)

	req := rating.GetRequest{Id: request.RatingId}

	response, err := ratingClient.Get(context.TODO(), &req)
	if err == nil {
		getAverageRatingByHostIdRequest := rating.GetAverageRatingByHostIdRequest{Id: response.Rating.HostId}

		ratingResponse, err := ratingClient.GetAverageRatingByHostId(context.TODO(), &getAverageRatingByHostIdRequest)
		if err == nil {

			userClient := services.NewUserClient(service.userClientAddress)
			userGetRequest := user.GetRequest{Id: response.Rating.HostId}
			host, err := userClient.Get(context.TODO(), &userGetRequest)
			if err == nil {

				if ratingResponse.AverageRating > 4.7 {
					if host.User.Outstanding != "YES" {
						host.User.Outstanding = "YES"
						userClient.UpdateUser(context.TODO(), &user.UpdateUserRequest{User: host.User})
					}
				} else {
					if host.User.Outstanding != "NO" {
						host.User.Outstanding = "NO"
						userClient.UpdateUser(context.TODO(), &user.UpdateUserRequest{User: host.User})
					}
				}
			}
		}
	}
	authDeleteRatingResposne := pb.AuthDeleteRatingResponse{}
	fmt.Print("authDeleteRatingResposne: ")
	fmt.Println(authDeleteRatingResposne)
	return &authDeleteRatingResposne, nil
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

	accRequest := accommodation.GetMessageHostReguest{Id: bookingResponse.Booking.AccommodationId}

	accommodationClient := services.NewAccommodationClient(service.accommodationClientAddress)

	accResponse, err := accommodationClient.GetHostId(context.TODO(), &accRequest)
	if err == nil {
		cancelingRequest := booking.GetCancellationRateForHostRequest{HostId: accResponse.HostId}

		ratingResponse, err := bookingClient.GetCancellationRateForHost(context.TODO(), &cancelingRequest)
		if err == nil {

			userClient := services.NewUserClient(service.userClientAddress)
			userGetRequest := user.GetRequest{Id: accResponse.HostId}
			host, err := userClient.Get(context.TODO(), &userGetRequest)
			if err == nil {

				if ratingResponse.Percentage < 5 {
					if host.User.Outstanding != "YES" {
						host.User.Outstanding = "YES"
						userClient.UpdateUser(context.TODO(), &user.UpdateUserRequest{User: host.User})
					}
				} else {
					if host.User.Outstanding != "NO" {
						host.User.Outstanding = "NO"
						userClient.UpdateUser(context.TODO(), &user.UpdateUserRequest{User: host.User})
					}
				}
			}
		}
	}

	fmt.Print("bookingResponse: ")
	fmt.Println(bookingResponse)
	authCanceling := pb.AuthReservationCancelingResponse{Booking: &pb.Booking{Id: bookingResponse.Booking.Id, AccommodationId: bookingResponse.Booking.AccommodationId, GuestId: bookingResponse.Booking.GuestId, Price: bookingResponse.Booking.Price, PriceType: pb.Booking_PriceType(bookingResponse.Booking.PriceType), NumberOfGuests: bookingResponse.Booking.NumberOfGuests, BookingType: pb.Booking_BookingType(bookingResponse.Booking.BookingType), StartDate: bookingResponse.Booking.StartDate, EndDate: bookingResponse.Booking.EndDate}}
	fmt.Print("authCreateBookingResponse: ")
	fmt.Println(authCanceling)
	return &authCanceling, nil
}

func (service *AuthService) GetAccommodationsByHostId(jwtData *domain.JwtData) (*pb.AuthGetAccommodationsByHostIdResponse, error) {
	getAccommodationsRequest := accommodation.GetByHostIdRequest{HostId: jwtData.UserId}
	accommodationClient := services.NewAccommodationClient(service.accommodationClientAddress)
	accommodationResponse, err := accommodationClient.GetByHostId(context.TODO(), &getAccommodationsRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("accommodationResponse: ")
	fmt.Println(accommodationResponse)
	var respAccommodations []*pb.Accommodation
	for _, accommodationIt := range accommodationResponse.Acccommodations {
		temp := pb.Accommodation{
			Id:        accommodationIt.Id,
			HostId:    accommodationIt.HostId,
			Name:      accommodationIt.Name,
			Location:  accommodationIt.Location,
			Benefits:  accommodationIt.Benefits,
			Pictures:  accommodationIt.Pictures,
			MinGuests: accommodationIt.MinGuests,
			MaxGuests: accommodationIt.MaxGuests,
			Price:     accommodationIt.Price,
			PriceType: pb.Accommodation_PriceType(accommodationIt.PriceType),
		}
		respAccommodations = append(respAccommodations, &temp)
	}
	authGetAccommodationsByHostIdResponse := pb.AuthGetAccommodationsByHostIdResponse{Accommodations: respAccommodations}
	fmt.Print("authGetAccommodationsByHostIdResponse: ")
	fmt.Println(authGetAccommodationsByHostIdResponse)
	return &authGetAccommodationsByHostIdResponse, nil
}

func (service *AuthService) GetUserRatingByAccommodationId(jwtData *domain.JwtData, request *pb.AuthGetUserRatingByAccommodationIdRequest) (*pb.AuthGetUserRatingByAccommodationIdResponse, error) {
	ratingRequest := rating.GetUserRatingByAccommodationIdRequest{AccommodationId: request.AccommodationId, GuestId: jwtData.UserId}
	ratingClient := services.NewRatingClient(service.ratingClientAddress)
	ratingResponse, err := ratingClient.GetUserRatingByAccommodationId(context.TODO(), &ratingRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("ratingResponse: ")
	fmt.Println(ratingResponse)

	authGetUserRatingByAccommodationIdResponse := pb.AuthGetUserRatingByAccommodationIdResponse{Rating: &pb.Rating{Id: ratingResponse.Rating.Id, HostId: ratingResponse.Rating.HostId, AccommodationId: ratingResponse.Rating.AccommodationId, GuestId: ratingResponse.Rating.GuestId, Rate: ratingResponse.Rating.Rate}}
	fmt.Print("authGetUserRatingByAccommodationIdResponse: ")
	fmt.Println(authGetUserRatingByAccommodationIdResponse)
	return &authGetUserRatingByAccommodationIdResponse, nil
}

func (service *AuthService) GetUserRatingByHostId(jwtData *domain.JwtData, request *pb.AuthGetUserRatingByHostIdRequest) (*pb.AuthGetUserRatingByHostIdResponse, error) {
	ratingRequest := rating.GetUserRatingByHostIdRequest{HostId: request.HostId, GuestId: jwtData.UserId}
	ratingClient := services.NewRatingClient(service.ratingClientAddress)
	ratingResponse, err := ratingClient.GetUserRatingByHostId(context.TODO(), &ratingRequest)
	if err != nil {
		return nil, err
	}
	fmt.Print("ratingResponse: ")
	fmt.Println(ratingResponse)

	authGetUserRatingByHostIdResponse := pb.AuthGetUserRatingByHostIdResponse{Rating: &pb.Rating{Id: ratingResponse.Rating.Id, HostId: ratingResponse.Rating.HostId, AccommodationId: ratingResponse.Rating.AccommodationId, GuestId: ratingResponse.Rating.GuestId, Rate: ratingResponse.Rating.Rate}}
	fmt.Print("authGetUserRatingByHostIdResponse: ")
	fmt.Println(authGetUserRatingByHostIdResponse)
	return &authGetUserRatingByHostIdResponse, nil
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
	authDefineCustomPriceResponse := pb.AuthDefineCustomPriceResponse{Accommodation: &pb.Accommodation{Id: accommodationResponse.Accommodation.Id, HostId: accommodationResponse.Accommodation.HostId, Name: accommodationResponse.Accommodation.Name, Location: accommodationResponse.Accommodation.Location, Benefits: accommodationResponse.Accommodation.Benefits, Pictures: accommodationResponse.Accommodation.Pictures, MinGuests: accommodationResponse.Accommodation.MinGuests, MaxGuests: accommodationResponse.Accommodation.MaxGuests}}
	fmt.Print("authDefineCustomPriceResponse: ")
	fmt.Println(authDefineCustomPriceResponse)
	return &authDefineCustomPriceResponse, nil
}

func (service *AuthService) GetBookingsByAccommodationId(jwtData *domain.JwtData, accommodationId string) (*pb.AuthGetBookingsByAccommodationIdResponse, error) {
	bookingClient := services.NewBookingClient(service.bookingClientAddress)

	bookingResponse, err := bookingClient.GetByAccommodationId(context.TODO(), &booking.GetByAccommodationIdRequest{AccommodationId: accommodationId})
	if err != nil {
		return nil, err
	}
	fmt.Print("bookingResponse: ")
	fmt.Println(bookingResponse)
	var respBookings []*pb.Booking
	for _, bookingIt := range bookingResponse.Bookings {
		temp := pb.Booking{
			Id:              bookingIt.Id,
			AccommodationId: bookingIt.AccommodationId,
			GuestId:         bookingIt.GuestId,
			Price:           bookingIt.Price,
			PriceType:       pb.Booking_PriceType(bookingIt.PriceType),
			NumberOfGuests:  bookingIt.NumberOfGuests,
			BookingType:     pb.Booking_BookingType(bookingIt.BookingType),
			StartDate:       bookingIt.StartDate,
			EndDate:         bookingIt.EndDate,
		}
		respBookings = append(respBookings, &temp)
	}
	authGetBookingByAccommodationIdResponse := pb.AuthGetBookingsByAccommodationIdResponse{Bookings: respBookings}
	return &authGetBookingByAccommodationIdResponse, nil
}

/*
func (service *AuthService) GetHost(request *pb.AuthGetUserRequest) (*pb.AuthGetUserResponse, error) {
	request := user.GetRequest{Id: request.Id}
	userClient := services.NewUserClient(service.userClientAddress)
	ratingResponse, err := userClient.Get(context.TODO(), &request)
	if err != nil {
		return nil, err
	}
	fmt.Print("ratingResponse: ")
	fmt.Println(ratingResponse)

	authGetUserRatingByHostIdResponse := pb.AuthGetUserRatingByHostIdResponse{Rating: &pb.Rating{Id: ratingResponse.Rating.Id, HostId: ratingResponse.Rating.HostId, AccommodationId: ratingResponse.Rating.AccommodationId, GuestId: ratingResponse.Rating.GuestId, Rate: ratingResponse.Rating.Rate}}
	fmt.Print("authGetUserRatingByHostIdResponse: ")
	fmt.Println(authGetUserRatingByHostIdResponse)
	return &authGetUserRatingByHostIdResponse, nil
}
*/
