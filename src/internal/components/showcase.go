package components

import (
	"encoding/json"
	"fmt"
	"service_api/internal/persist"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type ServiceShowcases struct {
	pc           *conf.Params
	showcases    persist.ShowcasesContext
	masterEntity persist.MasterEntityContext
	log          *log.Helper
}

func NewShowcasesService(
	pc *conf.Params,
	masterEntity persist.MasterEntityContext,
	logger log.Logger,
	showcases persist.ShowcasesContext) *ServiceShowcases {
	return &ServiceShowcases{
		pc:           pc,
		showcases:    showcases,
		masterEntity: masterEntity,
		log:          log.NewHelper(logger),
	}
}

func (s *ServiceShowcases) GetShowcase(req *pb.GetShowcaseRequest) (*pb.GetShowcaseReply, error) {
	showcases := &pb.GetShowcaseReply{}
	uuid := req.Uuid
	r := s.showcases.FindByUUID(uuid)
	if r == nil {
		return nil, errors.New(500, "UUID_CLEARED", "Не найдено значение с переданным UUID")
	}
	if len(r.PriceList) < 1 {
		return nil, errors.New(500, "PRICE_LIST_CLEARED", "Пустой список, передаваемый в Price List")
	}
	var arr []string
	err := json.Unmarshal([]byte(r.PriceList), &arr)
	if err != nil {
		return nil, errors.New(500, "PRICE_LIST_CLEARED", "Не корреткно указано значение Price List")
	}
	showcases.Name = r.Name
	showcases.Success = true
	for _, el := range arr {
		var n string
		_, err := fmt.Sscan(el, &n)
		if err != nil {
			continue
		}
		me := s.masterEntity.OneToPriceIdLarge(n)
		if me != nil &&
			me.Product != nil &&
			me.Price != nil &&
			me.Specification != nil && len(me.Specification) > 0 &&
			me.Images != nil {
			showcases.MasterEntity = append(showcases.MasterEntity, me)
		}
	}
	return showcases, nil
}
