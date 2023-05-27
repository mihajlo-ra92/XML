package api

import (
	"context"
	"fmt"

	"github.com/mihajlo-ra92/XML/rating_service/application"

	pb "github.com/mihajlo-ra92/XML/common/proto/rating_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingHandler struct {
	pb.UnimplementedRatingServiceServer
	service *application.RatingService
}

func NewRatingHandler(service *application.RatingService) *RatingHandler {
	return &RatingHandler{
		service: service,
	}
}

func (handler *RatingHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	rating, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	ratingPb := mapRating(rating)
	response := &pb.GetResponse{
		Rating: ratingPb,
	}
	return response, nil
}

func (handler *RatingHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	fmt.Println("In GetAll grpc api")
	ratings, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Ratings: []*pb.Rating{},
	}
	for _, rating := range ratings {
		current := mapRating(rating)
		response.Ratings = append(response.Ratings, current)
	}
	return response, nil
}

func (handler *RatingHandler) CreateRating(ctx context.Context, request *pb.CreateRatingRequest) (*pb.CreateRatingResponse, error) {
	fmt.Println("In CreateRating grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)
	rating := mapNewRating(request)
	fmt.Print("rating after mapping: ")
	fmt.Println(rating)
	err := handler.service.Create(rating)
	if err != nil {
		return nil, err
	}
	response := pb.CreateRatingResponse{Rating: mapRating(rating)}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *RatingHandler) DeleteRating(ctx context.Context, request *pb.DeleteRatingRequest) (*pb.DeleteRatingResponse, error) {
	fmt.Println("In DeleteRating grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)

	err := handler.service.Delete(request.RatingId)
	if err != nil {
		return nil, err
	}

	response := pb.DeleteRatingResponse{}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *RatingHandler) GetUserRatingByAccommodationId(ctx context.Context, request *pb.GetUserRatingByAccommodationIdRequest) (*pb.GetUserRatingByAccommodationIdResponse, error) {
	fmt.Println("In GetUserRatingByAccommodationId grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)
	rating, err := handler.service.GetUserRatingByAccommodationId(request.AccommodationId, request.GuestId)
	if err != nil {
		return nil, err
	}
	response := pb.GetUserRatingByAccommodationIdResponse{Rating: &pb.Rating{Id: rating.Id.Hex(), HostId: rating.HostId, AccommodationId: rating.AccommodationId, GuestId: rating.GuestId, Rate: rating.Rate}}

	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *RatingHandler) GetUserRatingByHostId(ctx context.Context, request *pb.GetUserRatingByHostIdRequest) (*pb.GetUserRatingByHostIdResponse, error) {
	fmt.Println("In GetUserRatingByHostId grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)
	rating, err := handler.service.GetUserRatingByHostId(request.HostId, request.GuestId)
	if err != nil {
		return nil, err
	}
	response := pb.GetUserRatingByHostIdResponse{Rating: &pb.Rating{Id: rating.Id.Hex(), HostId: rating.HostId, AccommodationId: rating.AccommodationId, GuestId: rating.GuestId, Rate: rating.Rate}}

	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}
