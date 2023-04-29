package api

import (
	"github.com/mihajlo-ra92/XML/accommodation_service/domain"
	pb "github.com/mihajlo-ra92/XML/common/proto/accommodation_service"
)

func mapAccommodation(accommodation *domain.Accommodation) *pb.Accommodation {
	accommodationPb := &pb.Accommodation{
		Id:		accommodation.Id.Hex(),
		Name:	accommodation.Name,
		Location:	accommodation.Location,
		Benefits: 		accommodation.Benefits,
		Pictures:	accommodation.Pictures,
		MinGuests:	accommodation.MinGuests,
		MaxGuests:	accommodation.MaxGuests,
	}
	return accommodationPb
}