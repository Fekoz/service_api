package components

import (
	"service_api/internal/persist"
	"service_api/src/util/consts"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type ServiceUtils struct {
	pc  *conf.Params
	uri persist.UriContext
	log *log.Helper
}

func NewUtilsService(
	pc *conf.Params,
	logger log.Logger,
	uri persist.UriContext) *ServiceUtils {
	return &ServiceUtils{
		pc:  pc,
		uri: uri,
		log: log.NewHelper(logger),
	}
}

func (s *ServiceUtils) GetMarketToUid(req *pb.GetMarketToUidRequest) (*pb.GetMarketToUidReply, error) {
	var item []string
	data := &pb.GetMarketToUidReply{
		Success: false,
	}

	for i, v := range req.Uid {
		if i > int(consts.MaxLimit) || i < 0 {
			break
		}
		item = append(item, v)
	}

	if len(item) < 1 {
		return nil, errors.New(500, "MARKET_URI", "Ошибка в получении данных UID")
	}

	result := s.uri.FindByMIDs(item)

	if len(result) > 0 {
		data.Success = true
	}

	for _, val := range result {
		data.Result = append(data.Result, &pb.Market{
			Uid: val.Uid,
			Url: val.Uri,
		})
	}

	return data, nil
}
