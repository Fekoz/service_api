package components

import (
	"service_api/internal/persist"
	"service_api/src/entitys"
	"service_api/src/util"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type ServiceCarpetOnly struct {
	pc            *conf.Params
	product       persist.ProductContext
	images        persist.ImagesContext
	price         persist.PriceContext
	specification persist.SpecificationContext
	attribute     persist.AttributeContext
	masterEntity  persist.MasterEntityContext
	collect       persist.CollectionContext
	log           *log.Helper
}

func NewCarpetOnlyService(
	pc *conf.Params,
	logger log.Logger,
	product persist.ProductContext,
	images persist.ImagesContext,
	price persist.PriceContext,
	specification persist.SpecificationContext,
	attribute persist.AttributeContext,
	collect persist.CollectionContext,
	masterEntity persist.MasterEntityContext) *ServiceCarpetOnly {
	return &ServiceCarpetOnly{
		pc:            pc,
		product:       product,
		images:        images,
		price:         price,
		specification: specification,
		attribute:     attribute,
		masterEntity:  masterEntity,
		collect:       collect,
		log:           log.NewHelper(logger),
	}
}

func (s *ServiceCarpetOnly) formatToPrice(item []entitys.Price) (*pb.GetCarpetOnlyPriceReply, error) {
	if item == nil || len(item) == 0 {
		return nil, errors.New(500, "CARPET_PRICE_ITEM_HAS_NULL", "Указан не верный параметр ARTICLE, UUID, ID")
	}
	var element []*pb.Price
	for _, el := range item {
		element = append(element, &pb.Price{
			Id:     int64(el.ID),
			Price:  int64(el.Price),
			Count:  el.Count,
			Height: el.Height,
			Width:  el.Width,
		})
	}
	return &pb.GetCarpetOnlyPriceReply{
		Success: true,
		Price:   element,
	}, nil
}

func (s *ServiceCarpetOnly) GetCarpetOnlyProduct(req *pb.GetCarpetOnlyProductRequest) (*pb.GetCarpetOnlyProductReply, error) {
	if len(req.Article) > 0 {
		toProductId := s.masterEntity.ArticleToProductId(req.Article)
		if toProductId == 0 {
			return nil, errors.New(500, "CARPET_PRODUCT_ITEM_HAS_NULL", "Ничего не найдено")
		}
		return util.FormatToProduct(s.product.OneToProductIdProduct(toProductId))
	}
	if len(req.Uuid) > 0 {
		toProductId := s.masterEntity.UuidToProductId(req.Uuid)
		if toProductId == 0 {
			return nil, errors.New(500, "CARPET_PRODUCT_ITEM_HAS_NULL", "Ничего не найдено")
		}
		return util.FormatToProduct(s.product.OneToProductIdProduct(toProductId))
	}
	if req.Id != 0 {
		return util.FormatToProduct(s.product.OneToProductIdProduct(req.Id))
	}
	return nil, errors.New(500, "CARPET_PRODUCT_ITEM_HAS_NULL", "Не указаны переданные значения ARTICLE, UUID, ID")
}

func (s *ServiceCarpetOnly) GetCarpetOnlyPrice(req *pb.GetCarpetOnlyPriceRequest) (*pb.GetCarpetOnlyPriceReply, error) {
	if len(req.Article) > 0 {
		toProductId := s.masterEntity.ArticleToProductId(req.Article)
		if toProductId == 0 {
			return nil, errors.New(500, "CARPET_PRICE_ITEM_HAS_NULL", "Ничего не найдено")
		}
		return s.formatToPrice(s.price.OneToProductIdPrice(toProductId))
	}
	if len(req.Uuid) > 0 {
		toProductId := s.masterEntity.UuidToProductId(req.Uuid)
		if toProductId == 0 {
			return nil, errors.New(500, "CARPET_PRICE_ITEM_HAS_NULL", "Ничего не найдено")
		}
		return s.formatToPrice(s.price.OneToProductIdPrice(toProductId))
	}
	if req.Id != 0 {
		return s.formatToPrice(s.price.OneToProductIdPrice(req.Id))
	}
	return nil, errors.New(500, "CARPET_PRICE_ITEM_HAS_NULL", "Не указаны переданные значения ARTICLE, UUID, ID")
}

func (s *ServiceCarpetOnly) GetCarpetOnlySpecification(req *pb.GetCarpetOnlySpecificationRequest) (*pb.GetCarpetOnlySpecificationReply, error) {
	if len(req.Article) > 0 {
		toProductId := s.masterEntity.ArticleToProductId(req.Article)
		if toProductId == 0 {
			return nil, errors.New(500, "CARPET_SPECIFICATION_ITEM_HAS_NULL", "Ничего не найдено")
		}
		return util.FormatToSpecification(s.specification.OneToProductIdSpecification(toProductId))
	}
	if len(req.Uuid) > 0 {
		toProductId := s.masterEntity.UuidToProductId(req.Uuid)
		if toProductId == 0 {
			return nil, errors.New(500, "CARPET_SPECIFICATION_ITEM_HAS_NULL", "Ничего не найдено")
		}
		return util.FormatToSpecification(s.specification.OneToProductIdSpecification(toProductId))
	}
	if req.Id != 0 {
		return util.FormatToSpecification(s.specification.OneToProductIdSpecification(req.Id))
	}
	return nil, errors.New(500, "CARPET_SPECIFICATION_ITEM_HAS_NULL", "Не указаны переданные значения ARTICLE, UUID, ID")
}

func (s *ServiceCarpetOnly) GetCarpetOnlyAttribute(req *pb.GetCarpetOnlyAttributeRequest) (*pb.GetCarpetOnlyAttributeReply, error) {
	if len(req.Article) > 0 {
		toProductId := s.masterEntity.ArticleToProductId(req.Article)
		if toProductId == 0 {
			return nil, errors.New(500, "CARPET_ATTRIBUTE_ITEM_HAS_NULL", "Ничего не найдено")
		}
		return util.FormatToAttribute(s.attribute.OneToProductIdAttribute(toProductId))
	}
	if len(req.Uuid) > 0 {
		toProductId := s.masterEntity.UuidToProductId(req.Uuid)
		if toProductId == 0 {
			return nil, errors.New(500, "CARPET_ATTRIBUTE_ITEM_HAS_NULL", "Ничего не найдено")
		}
		return util.FormatToAttribute(s.attribute.OneToProductIdAttribute(toProductId))
	}
	if req.Id != 0 {
		return util.FormatToAttribute(s.attribute.OneToProductIdAttribute(req.Id))
	}
	return nil, errors.New(500, "CARPET_ATTRIBUTE_ITEM_HAS_NULL", "Не указаны переданные значения ARTICLE, UUID, ID")
}

func (s *ServiceCarpetOnly) GetCarpetOnlyImages(req *pb.GetCarpetOnlyImagesRequest) (*pb.GetCarpetOnlyImagesReply, error) {
	if len(req.Article) > 0 {
		toProductId := s.masterEntity.ArticleToProductId(req.Article)
		if toProductId == 0 {
			return nil, errors.New(500, "CARPET_IMAGES_ITEM_HAS_NULL", "Ничего не найдено")
		}
		return util.FormatToImages(s.images.OneToProductIdImages(toProductId))
	}
	if len(req.Uuid) > 0 {
		toProductId := s.masterEntity.UuidToProductId(req.Uuid)
		if toProductId == 0 {
			return nil, errors.New(500, "CARPET_IMAGES_ITEM_HAS_NULL", "Ничего не найдено")
		}
		return util.FormatToImages(s.images.OneToProductIdImages(toProductId))
	}
	if req.Id != 0 {
		return util.FormatToImages(s.images.OneToProductIdImages(req.Id))
	}
	return nil, errors.New(500, "CARPET_IMAGES_ITEM_HAS_NULL", "Не указаны переданные значения ARTICLE, UUID, ID")
}

func (s *ServiceCarpetOnly) GetCarpetCollection(req *pb.GetCarpetCollectionRequest) (*pb.GetCarpetCollectionReply, error) {
	if len(req.Article) > 0 {
		toProductId := s.masterEntity.ArticleToProductId(req.Article)
		if toProductId == 0 {
			return nil, errors.New(500, "CARPET_COLLECTION_ITEM_HAS_NULL", "Ничего не найдено")
		}
		var collection []string
		collect := s.collect.FindWithId(toProductId)
		for _, c := range collect {
			data := s.masterEntity.ProductIdToArticle(c.OutProductId)
			if len(data) > 0 {
				collection = append(collection, data)
			}
		}
		return &pb.GetCarpetCollectionReply{
			Success:  true,
			Article:  req.Article,
			Articles: collection,
		}, nil
	}

	return nil, errors.New(500, "CARPET_COLLECTION_ITEM_HAS_NULL", "Не указаны переданные значения ARTICLE")
}
