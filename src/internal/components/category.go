package components

import (
	"service_api/internal/cache"
	"service_api/internal/persist"
	"service_api/src/util"
	"service_api/src/util/consts"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/log"
)

type ServiceCategory struct {
	pc       *conf.Params
	category persist.CategoryContext
	cache    *cache.Service
	log      *log.Helper
}

func NewCategoryService(
	pc *conf.Params,
	logger log.Logger,
	cache *cache.Service,
	category persist.CategoryContext) *ServiceCategory {
	return &ServiceCategory{
		pc:       pc,
		category: category,
		cache:    cache,
		log:      log.NewHelper(logger),
	}
}

func (s *ServiceCategory) supplyCategory() *pb.GetCatalogCategoryReply {
	var reply []*pb.GetCatalogCategoryReply_Type
	category := s.category.FindIsActive()
	for _, c := range category {
		//reply = util.AppendTypeCategorys  (reply, c.Type)
		//reply = util.AppendKeyCategorys   (reply, c.Type, c.Key, c.Name)
		//reply = util.AppendValueCategorys (reply, c.Type, c.Key, c.Value)
		reply = util.AppendArrayCategory(reply, consts.TypeCastType, c.Type, "", "")
		reply = util.AppendArrayCategory(reply, consts.TypeCastKey, c.Type, c.Key, c.Name)
		reply = util.AppendArrayCategory(reply, consts.TypeCastValue, c.Type, c.Key, c.Value)
	}
	return &pb.GetCatalogCategoryReply{
		Type: reply,
	}
}

func (s *ServiceCategory) GetCategory(req *pb.GetCatalogCategoryRequest) (*pb.GetCatalogCategoryReply, error) {
	cached, err := s.cache.GetCategorysCache(req)
	if err != nil || cached == nil {
		data := s.supplyCategory()
		s.cache.SetCategorysCache(req, data)
		return data, nil
	}
	return cached, nil
}
