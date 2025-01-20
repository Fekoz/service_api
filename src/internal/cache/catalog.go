package cache

import (
	"errors"
	"service_api/src/util"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
)

func (c *Service) GetCatalogsCache(req *pb.GetCatalogRequest) (*pb.GetCatalogReply, error) {
	key, err := util.GenKey[*pb.GetCatalogRequest](c.conf.Cache.Catalog.Name, &req)
	if err != nil {
		return &pb.GetCatalogReply{}, err
	}

	data, err := util.DecodeValue[*pb.GetCatalogReply](c.read(key))
	if err != nil {
		return &pb.GetCatalogReply{}, err
	}

	if data == nil {
		return &pb.GetCatalogReply{}, errors.New("GetCatalogsCache data is not found")
	}

	return data, nil
}

func (c *Service) SetCatalogsCache(req *pb.GetCatalogRequest, resp *pb.GetCatalogReply) {
	key, err := util.GenKey[*pb.GetCatalogRequest](c.conf.Cache.Catalog.Name, &req)
	if err != nil {
		return
	}

	data, err := util.EncodeValue[*pb.GetCatalogReply](&resp)
	if err != nil {
		return
	}

	if len(data) == 0 {
		return
	}
	c.append(key, data, c.conf.Cache.Catalog.Ttl)
}
