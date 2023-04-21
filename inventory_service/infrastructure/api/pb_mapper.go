package api

import (
	pb "github.com/mihajlo-ra92/XML/common/proto/inventory_service"
	"github.com/mihajlo-ra92/XML/inventory_service/domain"
)

func mapProduct(product *domain.Product) *pb.Product {
	productPb := &pb.Product{
		Id:        product.ProductId,
		ColorCode: product.ColorCode,
		Quantity:  int64(product.Quantity),
	}
	return productPb
}
