package cache

import (
	"errors"
	"service_api/src/util"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
)

func (c *Service) GetCategorysCache(req *pb.GetCatalogCategoryRequest) (*pb.GetCatalogCategoryReply, error) {
	key, err := util.GenKey[*pb.GetCatalogCategoryRequest](c.conf.Cache.Category.Name, &req)
	if err != nil {
		return &pb.GetCatalogCategoryReply{}, err
	}

	if !c.is(key) {
		return &pb.GetCatalogCategoryReply{}, errors.New("GetGategorysCache is not found or delete with TTL")
	}

	data, err := util.DecodeValue[*pb.GetCatalogCategoryReply](c.read(key))
	if err != nil {
		return &pb.GetCatalogCategoryReply{}, err
	}

	if data == nil {
		return &pb.GetCatalogCategoryReply{}, errors.New("GetGategorysCache data is not found")
	}

	return data, nil
}

func (c *Service) SetCategorysCache(req *pb.GetCatalogCategoryRequest, resp *pb.GetCatalogCategoryReply) {
	key, err := util.GenKey[*pb.GetCatalogCategoryRequest](c.conf.Cache.Category.Name, &req)
	if err != nil {
		return
	}

	data, err := util.EncodeValue[*pb.GetCatalogCategoryReply](&resp)
	if err != nil {
		return
	}

	if len(data) == 0 {
		return
	}
	c.append(key, data, c.conf.Cache.Category.Ttl)
}
