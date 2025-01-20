package components_admin

import (
	"service_api/internal/persist"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type AdminServicePriceDynamicExport struct {
	pc  *conf.Params
	log *log.Helper
	pde persist.PriceExportDynamicContext
}

func NewAdminPriceDynamicExportService(
	pc *conf.Params,
	logger log.Logger,
	pde persist.PriceExportDynamicContext) *AdminServicePriceDynamicExport {
	return &AdminServicePriceDynamicExport{
		pc:  pc,
		pde: pde,
		log: log.NewHelper(logger),
	}
}

func (s *AdminServicePriceDynamicExport) List(req *pb.GetPriceExportDynamicsRequest) (*pb.GetPriceExportDynamicsReply, error) {
	var array []*pb.PriceExportDynamic
	count := s.pde.FindListCount()

	if count < 1 {
		return nil, errors.New(500, "PDE_SERVICE", "Пусто")
	}

	offset := req.Point*req.Limit - req.Limit

	qb := s.pde.FindList(req.Limit, offset)
	if qb == nil {
		return nil, errors.New(500, "PDE_SERVICE", "Ошибка в получении списка")
	}

	for _, el := range qb {
		if el == nil {
			continue
		}

		array = append(array, &pb.PriceExportDynamic{
			Id:            int64(el.ID),
			MinWidth:      el.MinWidth,
			MaxWidth:      el.MaxWidth,
			MinHeight:     el.MinHeight,
			MaxHeight:     el.MaxHeight,
			PackageWidth:  el.PackageWidth,
			PackageHeight: el.PackageHeight,
			PackageDepth:  el.PackageDepth,
		})
	}

	return &pb.GetPriceExportDynamicsReply{
		Success:            true,
		Limit:              req.Limit,
		Point:              req.Point,
		Count:              count,
		PriceExportDynamic: array,
	}, nil
}

func (s *AdminServicePriceDynamicExport) createEntity(req *pb.SetPriceExportDynamicRequest) (*pb.SetPriceExportDynamicReply, error) {
	var reply = &pb.SetPriceExportDynamicReply{Success: true}
	qb := s.pde.Create(
		req.PriceExportDynamic.MinWidth,
		req.PriceExportDynamic.MaxWidth,
		req.PriceExportDynamic.MinHeight,
		req.PriceExportDynamic.MaxHeight,
		req.PriceExportDynamic.PackageDepth,
		req.PriceExportDynamic.PackageWidth,
		req.PriceExportDynamic.PackageHeight,
	)
	if qb == 0 {
		return nil, errors.New(500, "PDE_SERVICE", "Ошибка в создании записи")
	}

	return reply, nil
}

func (s *AdminServicePriceDynamicExport) Set(req *pb.SetPriceExportDynamicRequest) (*pb.SetPriceExportDynamicReply, error) {
	var reply = &pb.SetPriceExportDynamicReply{Success: true}

	if req.Id == 0 {
		return s.createEntity(req)
	}

	pde := s.pde.FindByID(req.Id)

	if pde != nil {
		pde.MinWidth = req.PriceExportDynamic.MinWidth
		pde.MaxWidth = req.PriceExportDynamic.MaxWidth
		pde.MinHeight = req.PriceExportDynamic.MinHeight
		pde.MaxHeight = req.PriceExportDynamic.MaxHeight
		pde.PackageDepth = req.PriceExportDynamic.PackageDepth
		pde.PackageWidth = req.PriceExportDynamic.PackageWidth
		pde.PackageHeight = req.PriceExportDynamic.PackageHeight
		s.pde.Update(pde)
		return reply, nil
	}

	return s.createEntity(req)
}

func (s *AdminServicePriceDynamicExport) Drop(req *pb.DropPriceExportDynamicRequest) (*pb.DropPriceExportDynamicReply, error) {
	pde := s.pde.FindByID(req.Id)

	if pde == nil {
		return nil, errors.New(500, "PDE_SERVICE", "Такой записи нет")
	}

	s.pde.Drop(pde)
	return &pb.DropPriceExportDynamicReply{
		Success: true,
	}, nil
}
