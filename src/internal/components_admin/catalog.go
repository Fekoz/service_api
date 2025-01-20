package components_admin

import (
	"regexp"
	"service_api/internal/persist"
	"service_api/src/entitys"
	"strings"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type AdminServiceCatalog struct {
	pc           *conf.Params
	masterEntity persist.MasterEntityAdminContext
	images       persist.ImagesContext
	price        persist.PriceContext
	log          *log.Helper
}

func NewAdminCatalogService(
	pc *conf.Params,
	masterEntity persist.MasterEntityAdminContext,
	images persist.ImagesContext,
	price persist.PriceContext,
	logger log.Logger) *AdminServiceCatalog {
	return &AdminServiceCatalog{
		pc:           pc,
		masterEntity: masterEntity,
		images:       images,
		price:        price,
		log:          log.NewHelper(logger),
	}
}

func (s *AdminServiceCatalog) GetCatalog(req *pb.GetCatalogRequest) (*pb.GetCatalogReply, error) {
	var me []*pb.MasterEntityCrop
	var count int64
	count = s.masterEntity.GetCarpetCountParams(req.Product, req.Specification, req.Price)

	if count < 1 {
		return nil, errors.New(500, "CATALOG_COUNT", "Пусто")
	}

	offset := req.Point*req.Limit - req.Limit

	qb := s.masterEntity.GetCarpetProductListParams(req.Limit, offset, req.Product, req.Specification, req.Price)
	if qb == nil {
		return nil, errors.New(500, "CATALOG_LIST", "Ошибка в получении списка товаров")
	}

	for _, el := range qb {
		if el == nil {
			continue
		}
		images := s.images.FindToImagesFirstAdmin(el.ID)
		priceRange := s.price.RangePriceListAdmin(el.ID)
		me = append(me, &pb.MasterEntityCrop{
			Product: &pb.Product{
				Id:                  el.ID,
				Article:             el.Article,
				Uuid:                el.Article,
				IsActive:            el.IsActive,
				Name:                el.Name,
				IsGlobalLock:        el.IsGlobalUpdateLock,
				IsProductLock:       el.IsProductLock,
				IsImagesLock:        el.IsImagesLock,
				IsPriceLock:         el.IsPriceLock,
				IsSpecificationLock: el.IsSpecificationLock,
				IsAttributeLock:     el.IsAttributeLock,
			},
			Images:     images,
			PriceRange: priceRange,
		})
	}

	return &pb.GetCatalogReply{
		Limit:        req.Limit,
		Point:        req.Point,
		Count:        count,
		MasterEntity: me,
	}, nil
}

func (s *AdminServiceCatalog) SerachWithParam(req *pb.SerachWithParamRequest) (*pb.SerachWithParamReply, error) {
	var me []*pb.MasterEntityCrop
	var qb []*entitys.Product

	if len(req.Mid) > 0 {
		parts := strings.Split(req.Mid, ".")
		mid := parts[len(parts)-1]
		qb = s.masterEntity.GetProductsByMid(mid)
	}

	if len(req.Url) > 0 {
		url := regexp.MustCompile(`/product/[^/]+`).FindString(req.Url)
		if len(url) > 0 {
			qb = s.masterEntity.GetProductsByUrl(url)
		}
	}

	if qb == nil {
		return nil, errors.New(500, "CATALOG_LIST", "Ошибка в получении списка товаров")
	}

	for _, el := range qb {
		if el == nil {
			continue
		}
		images := s.images.FindToImagesFirstAdmin(el.ID)
		priceRange := s.price.RangePriceListAdmin(el.ID)
		me = append(me, &pb.MasterEntityCrop{
			Product: &pb.Product{
				Id:                  el.ID,
				Article:             el.Article,
				Uuid:                el.Article,
				IsActive:            el.IsActive,
				Name:                el.Name,
				IsGlobalLock:        el.IsGlobalUpdateLock,
				IsProductLock:       el.IsProductLock,
				IsImagesLock:        el.IsImagesLock,
				IsPriceLock:         el.IsPriceLock,
				IsSpecificationLock: el.IsSpecificationLock,
				IsAttributeLock:     el.IsAttributeLock,
			},
			Images:     images,
			PriceRange: priceRange,
		})
	}

	success := false
	if len(me) > 0 {
		success = true
	}

	return &pb.SerachWithParamReply{
		Success:      success,
		MasterEntity: me,
	}, nil
}
