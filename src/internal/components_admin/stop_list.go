package components_admin

import (
	"service_api/internal/persist"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type AdminServiceStopList struct {
	pc       *conf.Params
	stopList persist.StopListContext
	log      *log.Helper
}

func NewAdminStopListService(
	pc *conf.Params,
	logger log.Logger,
	stopList persist.StopListContext) *AdminServiceStopList {
	return &AdminServiceStopList{
		pc:       pc,
		stopList: stopList,
		log:      log.NewHelper(logger),
	}
}

func (s *AdminServiceStopList) GetStopLists(req *pb.GetStopListsRequest) (*pb.GetStopListsReply, error) {
	var array []*pb.StopList
	var count int64
	count = s.stopList.FindListCount()

	if count < 1 {
		return nil, errors.New(500, "STOP_LIST", "Пусто")
	}

	offset := req.Point*req.Limit - req.Limit

	qb := s.stopList.FindList(req.Limit, offset)
	if qb == nil {
		return nil, errors.New(500, "STOP_LIST", "Ошибка в получении списка")
	}

	for _, el := range qb {
		if el == nil {
			continue
		}

		array = append(array, &pb.StopList{
			Id:         int64(el.ID),
			Name:       el.Name,
			Ip:         el.Ip,
			Session:    el.Session,
			Date:       el.UpdateAt.String(),
			Type:       el.Type,
			DateUnlock: el.UnlockAt.String(),
			Try:        el.Try,
		})
	}

	return &pb.GetStopListsReply{
		Success:  true,
		Limit:    req.Limit,
		Point:    req.Point,
		Count:    count,
		StopList: array,
	}, nil
}

func (s *AdminServiceStopList) DropStopList(req *pb.DropStopListRequest) (*pb.DropStopListReply, error) {
	isUpdate := s.stopList.DropByID(req.Id)
	if true == isUpdate {
		return &pb.DropStopListReply{Success: true}, nil
	}

	return nil, errors.New(500, "STOP_LIST", "Ошибка при обновлении")
}
