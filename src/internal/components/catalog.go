package components

import (
	"service_api/internal/cache"
	"service_api/internal/persist"
	"service_api/src/util/consts"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type ServiceCatalog struct {
	pc           *conf.Params
	masterEntity persist.MasterEntityContext
	images       persist.ImagesContext
	price        persist.PriceContext
	cache        *cache.Service
	log          *log.Helper
}

func NewCatalogService(
	pc *conf.Params,
	masterEntity persist.MasterEntityContext,
	images persist.ImagesContext,
	price persist.PriceContext,
	cache *cache.Service,
	logger log.Logger) *ServiceCatalog {
	return &ServiceCatalog{
		pc:           pc,
		masterEntity: masterEntity,
		images:       images,
		price:        price,
		cache:        cache,
		log:          log.NewHelper(logger),
	}
}

func (s *ServiceCatalog) supplyCatalog(req *pb.GetCatalogRequest) (*pb.GetCatalogReply, error) {
	var me []*pb.MasterEntityCrop
	var count int64
	count = s.masterEntity.GetCarpetCountParams(req.Product, req.Specification, req.Price, req.Size)

	if count < 1 {
		return nil, errors.New(500, "CATALOG_COUNT", "Пусто")
	}
	offset := req.Point*req.Limit - req.Limit
	qb := s.masterEntity.GetCarpetProductListParams(req.Limit, offset, req.Product, req.Specification, req.Price, req.Size)
	if qb == nil {
		return nil, errors.New(500, "CATALOG_LIST", "Ошибка в получении списка товаров")
	}
	for _, el := range qb {
		if el == nil {
			continue
		}
		images := s.images.FindToImagesFirst(el.ID)
		if images == nil {
			images = &pb.Images{
				Name: consts.DefaultImageFile,
				Dir:  consts.DefaultImageDir,
				Type: true,
			}
		}
		priceRange := s.price.RangePriceList(el.ID)
		if priceRange.Count == 0 {
			priceRange = &pb.MasterEntityCrop_PriceRange{
				PriceStart: 0,
				Count:      0,
				PriceEnd:   0,
			}
		}
		me = append(me, &pb.MasterEntityCrop{
			Product: &pb.Product{
				Id:       el.ID,
				Article:  el.Article,
				Uuid:     el.Article,
				IsActive: el.IsActive,
				Name:     el.Name,
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

func (s *ServiceCatalog) GetCatalog(req *pb.GetCatalogRequest) (*pb.GetCatalogReply, error) {
	cached, err := s.cache.GetCatalogsCache(req)
	if err != nil || cached == nil {
		data, err := s.supplyCatalog(req)
		if err != nil {
			return data, err
		}

		if len(data.MasterEntity) > 0 {
			s.cache.SetCatalogsCache(req, data)
		}

		return data, nil
	}
	return cached, nil
}
