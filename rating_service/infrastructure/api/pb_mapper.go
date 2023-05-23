package api

import (
	pb "github.com/mihajlo-ra92/XML/common/proto/rating_service"
	"github.com/mihajlo-ra92/XML/rating_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapNewRating(request *pb.CreateRatingRequest) *domain.Rating {
	booking := &domain.Rating{

		Id:              primitive.NewObjectID(),
		HostId:          request.Rating.HostId,
		AccommodationId: request.Rating.AccommodationId,
		GuestId:         request.Rating.GuestId,
		Rate:            request.Rating.Rate,
	}
	return booking
}

func mapRating(rating *domain.Rating) *pb.Rating {
	ratingPb := &pb.Rating{
		Id:              rating.Id.Hex(),
		HostId:          rating.HostId,
		AccommodationId: rating.AccommodationId,
		GuestId:         rating.GuestId,
		Rate:            rating.Rate,
	}
	return ratingPb
}
