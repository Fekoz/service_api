package util

import (
	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
)

func GetRequestNullable(req *pb.GetCatalogRequest) bool {
	if len(req.Product.Name) == 0 &&
		len(req.Product.Article) == 0 &&
		len(req.Specification) == 0 &&
		len(req.Price.Height) == 0 &&
		len(req.Price.Width) == 0 &&
		req.Price.PriceMin == 0 &&
		req.Price.PriceMax == 0 {
		return true
	}
	return false
}
