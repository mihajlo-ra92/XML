package api

import (
	"github.com/mihajlo-ra92/XML/catalogue_service/domain"
	pb "github.com/mihajlo-ra92/XML/common/proto/catalogue_service"
)

func mapProduct(product *domain.Product) *pb.Product {
	productPb := &pb.Product{
		Id:            product.Id.Hex(),
		Name:          product.Name,
		ClothingBrand: product.ClothingBrand,
	}
	for _, color := range product.Colors {
		productPb.Colors = append(productPb.Colors, &pb.Color{
			Code: color.Code,
			Name: color.Name,
		})
	}
	return productPb
}
