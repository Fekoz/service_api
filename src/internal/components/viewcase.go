package components

import (
	"encoding/json"
	"fmt"
	"service_api/internal/cache"
	"service_api/internal/persist"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type ServiceViewcases struct {
	pc           *conf.Params
	viewcases    persist.ViewcasesContext
	masterEntity persist.MasterEntityContext
	cache        *cache.Service
	log          *log.Helper
}

func NewViewcasesService(
	pc *conf.Params,
	masterEntity persist.MasterEntityContext,
	cache *cache.Service,
	logger log.Logger,
	viewcases persist.ViewcasesContext) *ServiceViewcases {
	return &ServiceViewcases{
		pc:           pc,
		viewcases:    viewcases,
		masterEntity: masterEntity,
		cache:        cache,
		log:          log.NewHelper(logger),
	}
}

func (s *ServiceViewcases) supplyViewcase(req *pb.GetViewcaseRequest) (*pb.GetViewcaseReply, error) {
	viewcases := &pb.GetViewcaseReply{}
	uuid := req.Uuid
	r := s.viewcases.FindByUUID(uuid)
	if r == nil {
		return nil, errors.New(500, "UUID_CLEARED", "Не найдено значение с переданным UUID")
	}
	if len(r.List) < 1 {
		return nil, errors.New(500, "LIST_CLEARED", "Пустой список, передаваемый в List")
	}
	var arr []string
	err := json.Unmarshal([]byte(r.List), &arr)
	if err != nil {
		return nil, errors.New(500, "LIST_CLEARED", "Не корреткно указано значение List")
	}
	viewcases.Name = r.Name
	viewcases.Success = true
	for _, el := range arr {
		var n string
		_, err := fmt.Sscan(el, &n)
		if err != nil {
			continue
		}
		me := s.masterEntity.OneToId(s.masterEntity.ArticleToProductId(n))
		if me != nil &&
			me.Product != nil &&
			me.Price != nil &&
			me.Specification != nil && len(me.Specification) > 0 &&
			me.Images != nil {
			viewcases.MasterEntity = append(viewcases.MasterEntity, me)
		}
	}
	return viewcases, nil
}

func (s *ServiceViewcases) GetViewcase(req *pb.GetViewcaseRequest) (*pb.GetViewcaseReply, error) {
	cached, err := s.cache.GetViewcaseCache(req)
	if err != nil || cached == nil {
		data, err := s.supplyViewcase(req)
		if err == nil {
			s.cache.SetViewcaseCache(req, data)
		}
		return data, err
	}
	return cached, nil
}
