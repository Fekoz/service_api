package components

import (
	"service_api/internal/cache"
	"service_api/internal/persist"
	"service_api/src/util"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type ServiceCarpet struct {
	pc           *conf.Params
	product      persist.ProductContext
	masterEntity persist.MasterEntityContext
	cache        *cache.Service
	log          *log.Helper
}

func NewCarpetService(
	pc *conf.Params,
	masterEntity persist.MasterEntityContext,
	logger log.Logger,
	cache *cache.Service,
	product persist.ProductContext) *ServiceCarpet {
	return &ServiceCarpet{
		pc:           pc,
		product:      product,
		masterEntity: masterEntity,
		cache:        cache,
		log:          log.NewHelper(logger),
	}
}

func (s *ServiceCarpet) getCarpetUuid(uuid string) (*pb.GetCarpetReply, error) {
	product := s.product.FindByUuid(uuid)
	if product == nil || len(product.Article) < 1 || len(product.Uuid) < 1 {
		return nil, errors.New(500, "CARPET_UUID_NULL", "Не найден ковер с указанным UUID")
	}
	me := util.ProductMapperCarpetPriceAppend(s.masterEntity.OneToIdIsProduct(util.ProductMapperCarpet(product)))
	if len(me.Product.Uuid) > 1 {
		return &pb.GetCarpetReply{
			Success:      true,
			MasterEntity: me,
		}, nil
	}
	return nil, errors.New(500, "CARPET_UUID_CLEARED", "Не найден ковер с переданным UUID")
}

func (s *ServiceCarpet) getCarpetArticle(article string) (*pb.GetCarpetReply, error) {
	product := s.product.FindByArticle(article)
	if product == nil || len(product.Article) < 1 || len(product.Uuid) < 1 {
		return nil, errors.New(500, "CARPET_ARTICLE_NULL", "Не найден ковер с указанным ARTICLE")
	}
	me := util.ProductMapperCarpetPriceAppend(s.masterEntity.OneToIdIsProduct(util.ProductMapperCarpet(product)))
	if len(me.Product.Uuid) > 1 {
		return &pb.GetCarpetReply{
			Success:      true,
			MasterEntity: me,
		}, nil
	}
	return nil, errors.New(500, "CARPET_ARTICLE_CLEARED", "Не найден ковер с переданным ARTICLE")
}

func (s *ServiceCarpet) getCarpetId(id int64) (*pb.GetCarpetReply, error) {
	me := util.ProductMapperCarpetPriceAppend(s.masterEntity.OneToId(id))
	if me != nil && me.Product != nil {
		return &pb.GetCarpetReply{
			Success:      true,
			MasterEntity: me,
		}, nil
	}
	return nil, errors.New(500, "CARPET_ID_CLEARED", "Не найден ковер с переданным ID")
}

func (s *ServiceCarpet) supplyCarpet(req *pb.GetCarpetRequest) (*pb.GetCarpetReply, error) {
	if len(req.Article) > 0 {
		return s.getCarpetArticle(req.Article)
	}
	if len(req.Uuid) > 0 {
		return s.getCarpetUuid(req.Uuid)
	}
	if req.Id != 0 {
		return s.getCarpetId(req.Id)
	}
	return nil, errors.New(500, "REQUEST_HAS_NULL", "Не указаны переданные значения ARTICLE, UUID, ID")
}

func (s *ServiceCarpet) GetCarpet(req *pb.GetCarpetRequest) (*pb.GetCarpetReply, error) {
	cached, err := s.cache.GetCarpetsCache(req)
	if err != nil || cached == nil {
		data, err := s.supplyCarpet(req)
		if err != nil {
			return data, err
		}

		if data.MasterEntity != nil {
			s.cache.SetCarpetsCache(req, data)
		}

		return data, nil
	}
	return cached, nil
}
