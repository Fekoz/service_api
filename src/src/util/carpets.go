package util

import (
	"service_api/src/entitys"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
)

func ProductMapperCarpets(product *entitys.Product) *pb.Product {
	return &pb.Product{
		Id:       product.ID,
		Article:  product.Article,
		Uuid:     product.Uuid,
		IsActive: product.IsActive,
		Name:     product.Name,
	}
}
