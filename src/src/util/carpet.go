package util

import (
	"service_api/src/entitys"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
)

func ProductMapperCarpet(product *entitys.Product) *pb.Product {
	return &pb.Product{
		Id:       product.ID,
		Article:  product.Article,
		Uuid:     product.Uuid,
		IsActive: product.IsActive,
		Name:     product.Name,
	}
}

func ProductMapperCarpetPriceAppend(p *pb.MasterEntity) *pb.MasterEntity {
	var list []*pb.Price

	for _, p := range p.Price {
		var isAppend bool
		isAppend = false
		for _, item := range list {
			if p.Mid == item.Mid || item.Price == p.Price {
				item.Count = item.Count + p.Count
				isAppend = true

				if p.Price > item.Price {
					item.Price = p.Price
				}

				if len(item.Uuid) > 0 {
					item.Uuids = append(item.Uuids, item.Uuid)
					item.Uuid = ""
				}

				if len(item.Uuids) > 0 {
					item.Uuids = append(item.Uuids, p.Uuid)
					continue
				}
			}
		}

		if !isAppend {
			list = append(list, &pb.Price{
				Id:     p.Id,
				Mid:    p.Mid,
				Price:  p.Price,
				Count:  p.Count,
				Height: p.Height,
				Width:  p.Width,
				Meter:  p.Meter,
				Uuid:   p.Uuid,
			})
		}
	}
	p.Price = list
	return p
}
