package api

import (
	"github.com/mihajlo-ra92/XML/accommodation_service/domain"
	pb "github.com/mihajlo-ra92/XML/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapAccommodation(accommodation *domain.Accommodation) *pb.Accommodation {
	accommodationPb := &pb.Accommodation{
		Id:        accommodation.Id.Hex(),
		HostId:    accommodation.HostId,
		Name:      accommodation.Name,
		Location:  accommodation.Location,
		Benefits:  accommodation.Benefits,
		Pictures:  accommodation.Pictures,
		MinGuests: accommodation.MinGuests,
		MaxGuests: accommodation.MaxGuests,
		Price:     accommodation.Price,
	}
	return accommodationPb
}

func mapNewAccommodation(request *pb.CreateAccommodationRequest) *domain.Accommodation {
	accommodation := &domain.Accommodation{
		Id:        primitive.NewObjectID(),
		HostId:    request.User.Id,
		Name:      request.Name,
		Location:  request.Location,
		Benefits:  request.Benefits,
		Pictures:  request.Pictures,
		MinGuests: request.MinGuests,
		MaxGuests: request.MaxGuests,
	}
	return accommodation
}
