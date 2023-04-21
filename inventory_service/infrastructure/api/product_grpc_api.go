package api

import (
	"context"

	pb "github.com/mihajlo-ra92/XML/common/proto/inventory_service"
	"github.com/mihajlo-ra92/XML/inventory_service/application"
)

type ProductHandler struct {
	service *application.ProductService
	pb.UnimplementedInventoryServiceServer
}

func NewProductHandler(service *application.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (handler *ProductHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	products, err := handler.service.GetAll()
	if err != nil || *products == nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Products: []*pb.Product{},
	}
	for _, product := range *products {
		current := mapProduct(&product)
		response.Products = append(response.Products, current)
	}
	return response, nil
}
