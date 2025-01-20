package cache

import (
	"errors"
	"service_api/src/util"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
)

func (c *Service) GetCarpetsCache(req *pb.GetCarpetRequest) (*pb.GetCarpetReply, error) {
	key, err := util.GenKey[*pb.GetCarpetRequest](c.conf.Cache.Carpet.Name, &req)
	if err != nil {
		return &pb.GetCarpetReply{}, err
	}

	data, err := util.DecodeValue[*pb.GetCarpetReply](c.read(key))
	if err != nil {
		return &pb.GetCarpetReply{}, err
	}

	if data == nil {
		return &pb.GetCarpetReply{}, errors.New("GetCarpetsCache data is not found")
	}

	return data, nil
}

func (c *Service) SetCarpetsCache(req *pb.GetCarpetRequest, resp *pb.GetCarpetReply) {
	key, err := util.GenKey[*pb.GetCarpetRequest](c.conf.Cache.Carpet.Name, &req)
	if err != nil {
		return
	}

	data, err := util.EncodeValue[*pb.GetCarpetReply](&resp)
	if err != nil {
		return
	}

	if len(data) == 0 {
		return
	}
	c.append(key, data, c.conf.Cache.Carpet.Ttl)
}
