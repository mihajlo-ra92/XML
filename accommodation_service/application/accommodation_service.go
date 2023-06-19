package application

import (
	"context"
	"fmt"

	"github.com/mihajlo-ra92/XML/accommodation_service/domain"
	"github.com/mihajlo-ra92/XML/accommodation_service/infrastructure/persistence"
	pb "github.com/mihajlo-ra92/XML/common/proto/accommodation_service"
	booking "github.com/mihajlo-ra92/XML/common/proto/booking_service"
	user "github.com/mihajlo-ra92/XML/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	store                domain.AccommodationStore
	bookingClientAddress string
	userClientAddress 	 string
}

func NewAccommodationService(store domain.AccommodationStore, bookingClientAddress string, userClientAddress string) *AccommodationService {
	return &AccommodationService{
		store:                store,
		bookingClientAddress: bookingClientAddress,
		userClientAddress: userClientAddress, 
	}
}

func (service *AccommodationService) Get(id primitive.ObjectID) (*domain.Accommodation, error) {
	return service.store.Get(id)
}

func (service *AccommodationService) GetByHostId(hostId string) ([]*domain.Accommodation, error) {
	return service.store.GetByHostId(hostId)
}

func (service *AccommodationService) GetAll() ([]*domain.Accommodation, error) {
	return service.store.GetAll()
}

func (service *AccommodationService) Create(accommodation *domain.Accommodation) error {
	return service.store.Insert(accommodation)
}

func (service *AccommodationService) DeleteAccommodationsByHostId(hostId string) error {
	accommodations, err := service.store.GetByHostId(hostId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Print("accommodations for host: " + hostId)
	fmt.Println(accommodations)
	bookingClient := persistence.NewBookingClient(service.bookingClientAddress)
	for _, accommodationIt := range accommodations{
		bookingResponse, err := bookingClient.DeleteBookingsByAccommodationId(context.TODO(), &booking.DeleteBookingByAccommodationIdRequest{AccommodationId: accommodationIt.HostId})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("bookingResponse: ")
		fmt.Println(bookingResponse)
	}
	err = service.store.DeleteByHostId(hostId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (service *AccommodationService) DefineCustomPrice(request pb.DefineCustomPriceRequest) error {
	fmt.Println("In DefineCustomPrice accommodation_service")
	bookingClient := persistence.NewBookingClient(service.bookingClientAddress)
	fmt.Println("booking_free_accomodation_search:")
	newBooking := booking.Booking{
		Id: "", 
		AccommodationId: request.AccommodationId, 
		GuestId: "", 
		Price: request.Price, 
		PriceType: booking.Booking_PriceType(request.PriceType), 
		NumberOfGuests: 0, 
		BookingType: booking.Booking_CustomPrice, 
		StartDate: request.StartDate, 
		EndDate: request.EndDate}
	bookingResponse, err := bookingClient.CreateBooking(context.TODO(), &booking.CreateBookingRequest{Booking: &newBooking})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Print("bookingResponse: ")
	fmt.Println(bookingResponse)
	return nil
}

func (service *AccommodationService) SearchWithFilter(request *pb.SearchRequest)([]*domain.Accommodation, error){
	
	fmt.Println("In Search accommodation_service")
	bookingClient := persistence.NewBookingClient(service.bookingClientAddress)
	userClient := persistence.NewUserClient(service.userClientAddress)
	fmt.Println("request:")
	fmt.Println(request)
	fmt.Println("booking_free_accomodation_search:")
	accommodations, err := service.store.SearchWithFilter(request.Location, request.Guest, request.MinPrice, request.MaxPrice, request.Benefits)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if (request.StartDate == nil) && (request.EndDate == nil) {
		return accommodations, nil
	}
	response := []*domain.Accommodation{}
		for _, accommodation := range accommodations {
			fmt.Println("accommodation id je  ", accommodation.Id.Hex())
			bookingResponse, err := bookingClient.GetByAccomodationIdandDataRange(context.TODO(), &booking.GetByAccomodationIdandDataRangeRequest{Id: accommodation.Id.Hex(), StartDate: request.StartDate, EndDate: request.EndDate})
			fmt.Println(bookingResponse)
			
			fmt.Println("udje ovde da se vidi booking")
			
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			var a uint32
			a = 0
			fmt.Println("da vidimo listu ; ", bookingResponse.Bookings)
			
			for _, booking := range bookingResponse.Bookings {
				fmt.Println("udje ovde da se vidi booking ; ", booking)
				
				a = a + booking.NumberOfGuests
				fmt.Println("sabiramo a")
				
			}
			
			fmt.Println("vrednost a : ", a)
			fmt.Println(accommodation.MaxGuests - request.Guest)
			
			if (accommodation.MaxGuests - request.Guest) >= a {
				if request.IsOutstanding == true{
					fmt.Println("ISOUTSTANDING JE TRUE")
					fmt.Println("ID U ISPITIVANJU OUTSTANDINGA JE :")
					fmt.Println(accommodation.HostId)
					userResponse, err := userClient.CheckIsOutstandingById(context.TODO(),&user.CheckIsOutstandingByIdRequest{Id : accommodation.HostId})
					if err != nil {
						return nil, err
					}else{
						if userResponse.IsOutstanding == true{
							response = append(response, accommodation)
						}
					}
						}else{
					response = append(response, accommodation)
				}
			}
			
		}
		fmt.Println("Work here : ")
		fmt.Println("booking search for free accomodation Response: ")
		fmt.Println("create search accommodation response: ")
	return response, nil


}

func (service *AccommodationService) Search(request *pb.SearchRequest) ([]*domain.Accommodation, error) {
	fmt.Println("In Search accommodation_service")
	bookingClient := persistence.NewBookingClient(service.bookingClientAddress)
	fmt.Println("request:")
	fmt.Println(request)
	fmt.Println("booking_free_accomodation_search:")
	accommodations, err := service.store.Search(request.Location, request.Guest)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if (request.StartDate == nil) && (request.EndDate == nil) {
		return accommodations, nil
	}
	response := []*domain.Accommodation{}
	for _, accommodation := range accommodations {
		fmt.Println("accommodation id je  ", accommodation.Id.Hex())
		bookingResponse, err := bookingClient.GetByAccomodationIdandDataRange(context.TODO(), &booking.GetByAccomodationIdandDataRangeRequest{Id: accommodation.Id.Hex(), StartDate: request.StartDate, EndDate: request.EndDate})
		fmt.Println(bookingResponse)

		fmt.Println("udje ovde da se vidi booking")

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		var a uint32
		a = 0
		fmt.Println("da vidimo listu ; ", bookingResponse.Bookings)

		for _, booking := range bookingResponse.Bookings {
			fmt.Println("udje ovde da se vidi booking ; ", booking)

			a = a + booking.NumberOfGuests
			fmt.Println("sabiramo a")

		}

		fmt.Println("vrednost a : ", a)
		fmt.Println(accommodation.MaxGuests - request.Guest)

		if (accommodation.MaxGuests - request.Guest) >= a {
			response = append(response, accommodation)
		}

	}

	fmt.Println("Work here : ")
	fmt.Println("booking search for free accomodation Response: ")
	fmt.Println("create search accommodation response: ")
	return response, nil

}

