package components_admin

import (
	"service_api/internal/persist"
	"service_api/src/entitys"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type AdminServiceCarpetOnly struct {
	pc           *conf.Params
	product      persist.ProductContext
	images       persist.ImagesContext
	price        persist.PriceContext
	masterEntity persist.MasterEntityContext
	log          *log.Helper
}

func NewAdminCarpetOnlyService(
	pc *conf.Params,
	logger log.Logger,
	product persist.ProductContext,
	images persist.ImagesContext,
	price persist.PriceContext,
	masterEntity persist.MasterEntityContext) *AdminServiceCarpetOnly {
	return &AdminServiceCarpetOnly{
		pc:           pc,
		product:      product,
		images:       images,
		price:        price,
		masterEntity: masterEntity,
		log:          log.NewHelper(logger),
	}
}

func (s *AdminServiceCarpetOnly) formatToPrice(item []entitys.Price) (*pb.GetCarpetOnlyPriceReply, error) {
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
			Mid:    el.Mid,
			Uuid:   el.Uuid,
			Meter:  el.Meter,
		})
	}
	return &pb.GetCarpetOnlyPriceReply{
		Success: true,
		Price:   element,
	}, nil
}

func (s *AdminServiceCarpetOnly) GetCarpetOnlyPrice(req *pb.GetCarpetOnlyPriceRequest) (*pb.GetCarpetOnlyPriceReply, error) {
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

func (s *AdminServiceCarpetOnly) GetCarpetMinToPriceUUID(req *pb.GetCarpetMinToPriceUUIDRequest) (*pb.GetCarpetMinToPriceUUIDReply, error) {
	if len(req.Uuid) < 1 {
		return nil, errors.New(500, "CARPET_UUID_PRICE_ITEM_HAS_NULL", "Пусто")
	}

	price := s.price.FindByUUID(req.Uuid)
	if price == nil || len(price.Uuid) < 1 {
		return nil, errors.New(500, "CARPET_UUID_PRICE_ITEM_HAS_NULL", "Ничего не найдено")
	}

	product := s.product.FindByID(price.ProductId)
	if product == nil {
		return nil, errors.New(500, "CARPET_UUID_PRICE_ITEM_HAS_NULL", "Ничего не найдено")
	}

	image := s.images.FindToImagesFirstAdmin(price.ProductId)
	if image == nil {
		return nil, errors.New(500, "CARPET_UUID_PRICE_ITEM_HAS_NULL", "Ничего не найдено")
	}

	return &pb.GetCarpetMinToPriceUUIDReply{
		Success: true,
		Article: product.Article,
		Images:  image,
		Id:      product.ID,
	}, nil
}

func (s *AdminServiceCarpetOnly) GetCarpetMinToPriceMID(req *pb.GetCarpetMinToPriceMIDRequest) (*pb.GetCarpetMinToPriceMIDReply, error) {
	if len(req.Mid) < 1 {
		return nil, errors.New(500, "CARPET_MID_PRICE_ITEM_HAS_NULL", "Пусто")
	}

	price := s.price.FindByMID(req.Mid)
	if price == nil || len(price.Uuid) < 1 {
		return nil, errors.New(500, "CARPET_MID_PRICE_ITEM_HAS_NULL", "Ничего не найдено")
	}

	product := s.product.FindByID(price.ProductId)
	if product == nil {
		return nil, errors.New(500, "CARPET_UUID_PRICE_ITEM_HAS_NULL", "Ничего не найдено")
	}

	image := s.images.FindToImagesFirstAdmin(price.ProductId)
	if image == nil {
		return nil, errors.New(500, "CARPET_UUID_PRICE_ITEM_HAS_NULL", "Ничего не найдено")
	}

	return &pb.GetCarpetMinToPriceMIDReply{
		Success: true,
		Article: product.Article,
		Images:  image,
		Id:      product.ID,
	}, nil
}

func (s *AdminServiceCarpetOnly) GetCarpetMinToArticle(req *pb.GetCarpetMinToArticleRequest) (*pb.GetCarpetMinToArticleReply, error) {
	if len(req.Article) < 1 {
		return nil, errors.New(500, "CARPET_ARTICLE_ITEM_HAS_NULL", "Пусто")
	}

	product := s.product.FindByArticle(req.Article)
	if product == nil {
		return nil, errors.New(500, "CARPET_ARTICLE_ITEM_HAS_NULL", "Ничего не найдено")
	}

	image := s.images.FindToImagesFirstAdmin(product.ID)
	if image == nil {
		return nil, errors.New(500, "CARPET_UUID_PRICE_ITEM_HAS_NULL", "Ничего не найдено")
	}

	return &pb.GetCarpetMinToArticleReply{
		Success: true,
		Article: product.Article,
		Images:  image,
		Id:      product.ID,
	}, nil
}
