package components_admin

import (
	"service_api/internal/persist"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type AdminServiceMarketSequence struct {
	pc             *conf.Params
	marketSequence persist.MarketSequenceContext
	price          persist.PriceContext
	log            *log.Helper
}

func NewAdminMarketSequenceService(
	pc *conf.Params,
	logger log.Logger,
	price persist.PriceContext,
	marketSequence persist.MarketSequenceContext) *AdminServiceMarketSequence {
	return &AdminServiceMarketSequence{
		pc:             pc,
		marketSequence: marketSequence,
		price:          price,
		log:            log.NewHelper(logger),
	}
}

func (s *AdminServiceMarketSequence) GetMarketSequences(req *pb.GetMarketSequencesRequest) (*pb.GetMarketSequencesReply, error) {
	var array []*pb.GetMarketSequencesReply_MarketSequence
	var count int64
	count = s.marketSequence.FindListCount()

	if count < 1 {
		return nil, errors.New(500, "MARKET_SEQUENCE", "Пусто")
	}

	offset := req.Point*req.Limit - req.Limit

	qb := s.marketSequence.FindList(req.Limit, offset)
	if qb == nil {
		return nil, errors.New(500, "MARKET_SEQUENCE", "Ошибка в получении списка")
	}

	for _, el := range qb {
		if el == nil {
			continue
		}
		array = append(array, &pb.GetMarketSequencesReply_MarketSequence{
			Id:         int64(el.ID),
			IsActive:   el.IsActive,
			Mid:        el.Mid,
			Name:       el.Name,
			Date:       el.UpdateAt.String(),
			IsCounter:  el.IsCounter,
			CounterPkg: el.CounterPkg,
		})
	}

	return &pb.GetMarketSequencesReply{
		Success:        true,
		Limit:          req.Limit,
		Point:          req.Point,
		Count:          count,
		MarketSequence: array,
	}, nil
}

func (s *AdminServiceMarketSequence) SetMarketSequence(req *pb.SetMarketSequenceRequest) (*pb.SetMarketSequenceReply, error) {
	mid := s.price.FindByMID(req.Mid)
	if len(mid.Mid) < 1 {
		return nil, errors.New(500, "MARKET_SEQUENCE", "Такого MID не существует")
	}
	s.marketSequence.SetActiveByMid(req.Mid, req.IsActive, req.IsCounter, req.CounterPkg)
	return &pb.SetMarketSequenceReply{
		Success: true,
	}, nil
}
