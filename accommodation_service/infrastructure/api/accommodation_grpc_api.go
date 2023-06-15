package api

import (
	"context"
	"fmt"

	"github.com/mihajlo-ra92/XML/accommodation_service/application"

	pb "github.com/mihajlo-ra92/XML/common/proto/accommodation_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	service *application.AccommodationService
}

func NewAccommodationHandler(service *application.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		service: service,
	}
}

func (handler *AccommodationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	accommodation, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	accommodationPb := mapAccommodation(accommodation)
	response := &pb.GetResponse{
		Accomodation: accommodationPb,
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	fmt.Println("In GetAll grpc api")
	accommodations, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Accommodations: []*pb.Accommodation{},
	}
	for _, accommodation := range accommodations {
		current := mapAccommodation(accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}
	return response, nil
}

func (handler *AccommodationHandler) CreateAccommodation(ctx context.Context, request *pb.CreateAccommodationRequest) (*pb.CreateAccommodationResponse, error) {
	fmt.Println("In CreateAccommodation grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)
	accommodation := mapNewAccommodation(request)
	fmt.Print("accommodation after mapping: ")
	fmt.Println(accommodation)
	err := handler.service.Create(accommodation)
	if err != nil {
		return nil, err
	}
	response := pb.CreateAccommodationResponse{Accommodation: mapAccommodation(accommodation)}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *AccommodationHandler) DeleteAccommodationsByHostId(ctx context.Context, request *pb.DeleteAccommodationsByHostIdRequest) (*pb.DeleteAccommodationsByHostIdResponse, error) {
	fmt.Println("In DeleteAccommodationsByHostId grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)
	err := handler.service.DeleteAccommodationsByHostId(request.HostId)
	if err != nil {
		return nil, err
	}
	response := pb.DeleteAccommodationsByHostIdResponse{Message: "Accommodations deleted"}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *AccommodationHandler) Search(ctx context.Context, request *pb.SearchRequest) (*pb.SearchResponse, error) {
	fmt.Println("InSearch grpc api")
	accommodations, err := handler.service.Search(request)
	if err != nil {
		return nil, err
	}

	response := &pb.SearchResponse{
		Accommodations: []*pb.AccommodationWithPrice{},
	}
	var accommodationWithPrice pb.AccommodationWithPrice

	for _, accommodation := range accommodations {
		current := mapAccommodation(accommodation)
		accommodationWithPrice.Accommodation = current
		accommodationWithPrice.Price = current.Price * request.Guest
		response.Accommodations = append(response.Accommodations, &accommodationWithPrice)
	}
	return response, nil
}

func (handler *AccommodationHandler) GetByHostId(ctx context.Context, request *pb.GetByHostIdRequest) (*pb.GetByHostIdResponse, error) {
	fmt.Println("In GetByHostId grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)
	accommodations, err := handler.service.GetByHostId(request.HostId)
	if err != nil {
		return nil, err
	}
	response := pb.GetByHostIdResponse{Acccommodations: []*pb.Accommodation{}}
	for _, accommodation := range accommodations {
		current := mapAccommodation(accommodation)
		response.Acccommodations = append(response.Acccommodations, current)
	}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *AccommodationHandler) DefineCustomPrice(ctx context.Context, request *pb.DefineCustomPriceRequest) (*pb.DefineCustomPriceResponse, error) {
	fmt.Println("In DefineCustomPriceRequest grpc api")
	fmt.Print("Request: ")
	fmt.Println(request)
	objectId, err := primitive.ObjectIDFromHex(request.AccommodationId)
	if err != nil {
		return nil, err
	}
	accommodation, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	if accommodation.HostId != request.User.Id {
		return nil, fmt.Errorf("user is not host of sent accommodation")
	}
	_ = handler.service.DefineCustomPrice(*request)
	response := pb.DefineCustomPriceResponse{Accommodation: mapAccommodation(accommodation)}
	fmt.Print("response: ")
	fmt.Println(response)
	return &response, nil
}

func (handler *AccommodationHandler) GetHostId(ctx context.Context, request *pb.GetMessageHostReguest) (*pb.GetMessageHostResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	accommodation, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	accommodationPb := mapAccommodation(accommodation)
	response := &pb.GetMessageHostResponse{
		HostId: accommodationPb.HostId,
	}
	return response, nil
}
