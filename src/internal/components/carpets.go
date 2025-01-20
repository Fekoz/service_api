package components

import (
	"service_api/internal/persist"
	"service_api/src/util"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type ServiceCarpets struct {
	pc           *conf.Params
	masterEntity persist.MasterEntityContext
	product      persist.ProductContext
	images       persist.ImagesContext
	price        persist.PriceContext
	log          *log.Helper
}

func NewCarpetsService(
	pc *conf.Params,
	masterEntity persist.MasterEntityContext,
	product persist.ProductContext,
	images persist.ImagesContext,
	price persist.PriceContext,
	logger log.Logger) *ServiceCarpets {
	return &ServiceCarpets{
		pc:           pc,
		masterEntity: masterEntity,
		product:      product,
		images:       images,
		price:        price,
		log:          log.NewHelper(logger),
	}
}

func (s *ServiceCarpets) getCarpetUuid(uuid string) *pb.MasterEntity {
	product := s.product.FindByUuid(uuid)
	if product == nil || len(product.Article) < 1 || len(product.Uuid) < 1 {
		return nil
	}
	me := s.masterEntity.OneToIdIsProduct(util.ProductMapperCarpets(product))
	if len(me.Product.Uuid) > 1 {
		return me
	}
	return nil
}

func (s *ServiceCarpets) getCarpetArticle(article string) *pb.MasterEntity {
	product := s.product.FindByArticle(article)
	if product == nil || len(product.Article) < 1 || len(product.Uuid) < 1 {
		return nil
	}
	me := s.masterEntity.OneToIdIsProduct(util.ProductMapperCarpets(product))
	if len(me.Product.Uuid) > 1 {
		return me
	}
	return nil
}

func (s *ServiceCarpets) getCarpetId(id int64) *pb.MasterEntity {
	me := s.masterEntity.OneToId(id)
	if me != nil && me.Product != nil {
		return me
	}
	return nil
}

func appendCarpetToArray(me []*pb.MasterEntity, item *pb.MasterEntity) []*pb.MasterEntity {
	if item != nil {
		me = append(me, &pb.MasterEntity{
			Product:       item.Product,
			Images:        item.Images,
			Price:         item.Price,
			Specification: item.Specification,
			Attribute:     item.Attribute,
		})
		return me
	}
	return nil
}

func (s *ServiceCarpets) GetCarpets(req *pb.GetCarpetsRequest) (*pb.GetCarpetsReply, error) {
	var me []*pb.MasterEntity
	if req.Products == nil || len(req.Products) == 0 {
		return nil, errors.New(500, "CARPETS_LIST", "Ошибка в получении списка товаров")
	}

	for idx, el := range req.Products {
		if me != nil && int64(len(me)) >= req.Limit {
			break
		}

		if int64(idx) <= (req.Limit*req.Point)-req.Limit && (idx >= 0 && req.Limit > 1) {
			continue
		}

		if int64(idx) >= (req.Limit*req.Point)+req.Limit {
			break
		}

		if len(el.Article) > 0 {
			me = appendCarpetToArray(me, s.getCarpetArticle(el.Article))
			continue
		}
		if len(el.Uuid) > 0 {
			me = appendCarpetToArray(me, s.getCarpetUuid(el.Article))
			continue
		}
		if el.Id != 0 {
			me = appendCarpetToArray(me, s.getCarpetId(el.Id))
			continue
		}
	}

	return &pb.GetCarpetsReply{
		Success:      true,
		Count:        int64(len(me)),
		MasterEntity: me,
		Point:        req.Point,
		Limit:        req.Limit,
	}, nil
}
