package cache

import (
	"errors"
	"service_api/src/util"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
)

func (c *Service) GetViewcaseCache(req *pb.GetViewcaseRequest) (*pb.GetViewcaseReply, error) {
	key, err := util.GenKey[*pb.GetViewcaseRequest](c.conf.Cache.Category.Name, &req)
	if err != nil {
		return &pb.GetViewcaseReply{}, err
	}

	if !c.is(key) {
		return &pb.GetViewcaseReply{}, errors.New("GetViewcaseReply is not found or delete with TTL")
	}

	data, err := util.DecodeValue[*pb.GetViewcaseReply](c.read(key))
	if err != nil {
		return &pb.GetViewcaseReply{}, err
	}

	if data == nil {
		return &pb.GetViewcaseReply{}, errors.New("GetViewcaseReply data is not found")
	}

	return data, nil
}

func (c *Service) SetViewcaseCache(req *pb.GetViewcaseRequest, resp *pb.GetViewcaseReply) {
	key, err := util.GenKey[*pb.GetViewcaseRequest](c.conf.Cache.Category.Name, &req)
	if err != nil {
		return
	}

	data, err := util.EncodeValue[*pb.GetViewcaseReply](&resp)
	if err != nil {
		return
	}

	if len(data) == 0 {
		return
	}
	c.append(key, data, c.conf.Cache.Category.Ttl)
}
