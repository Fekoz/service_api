package components_admin

import (
	"service_api/internal/persist"
	"service_api/src/entitys"
	"service_api/src/util/consts"

	conf "bitbucket.org/klaraeng/package_config/conf"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type AdminServiceCarpet struct {
	pc            *conf.Params
	product       persist.ProductContext
	price         persist.PriceContext
	specification persist.SpecificationContext
	images        persist.ImagesContext
	attribute     persist.AttributeContext
	masterEntity  persist.MasterEntityAdminContext
	log           *log.Helper
}

func NewAdminCarpetService(
	pc *conf.Params,
	masterEntity persist.MasterEntityAdminContext,
	logger log.Logger,
	product persist.ProductContext,
	price persist.PriceContext,
	specification persist.SpecificationContext,
	images persist.ImagesContext,
	attribute persist.AttributeContext,
) *AdminServiceCarpet {
	return &AdminServiceCarpet{
		pc:            pc,
		product:       product,
		price:         price,
		specification: specification,
		images:        images,
		attribute:     attribute,
		masterEntity:  masterEntity,
		log:           log.NewHelper(logger),
	}
}

func (s *AdminServiceCarpet) GetCarpet(req *pb.GetCarpetRequest) (*pb.GetCarpetReply, error) {
	if len(req.Article) < 1 {
		return nil, errors.New(500, "REQUEST_HAS_NULL", "Не указаны переданные значения ARTICLE")
	}

	product := s.product.FindByArticle(req.Article)
	if product == nil || len(product.Article) < 1 || len(product.Uuid) < 1 {
		return nil, errors.New(500, "CARPET_ARTICLE_NULL", "Не найден ковер с указанным ARTICLE")
	}

	me := s.masterEntity.OneToIdIsProduct(&pb.Product{
		Id:                  product.ID,
		Article:             product.Article,
		Uuid:                product.Uuid,
		IsActive:            product.IsActive,
		Name:                product.Name,
		IsGlobalLock:        product.IsGlobalUpdateLock,
		IsProductLock:       product.IsProductLock,
		IsPriceLock:         product.IsPriceLock,
		IsSpecificationLock: product.IsSpecificationLock,
		IsImagesLock:        product.IsImagesLock,
		IsAttributeLock:     product.IsAttributeLock,
		OriginalUrl:         product.OriginalUrl,
		DateUpdate:          product.UpdateAt.String(),
	})

	if len(me.Product.Uuid) > 1 {
		return &pb.GetCarpetReply{
			Success:      true,
			MasterEntity: me,
		}, nil
	}

	return nil, errors.New(500, "CARPET_ARTICLE_CLEARED", "Не найден ковер с переданным ARTICLE")
}

func (s *AdminServiceCarpet) SetCarpet(req *pb.SetCarpetRequest) (*pb.SetCarpetReply, error) {
	if len(req.Article) < 1 {
		return nil, errors.New(500, "REQUEST_HAS_NULL", "Не указаны переданные значения ARTICLE")
	}
	product := s.product.FindByArticle(req.Article)

	var productId int64
	if product == nil || product.ID == 0 {
		productId = s.product.Create(&entitys.Product{
			Article:             req.MasterEntity.Product.Article,
			Name:                req.MasterEntity.Product.Name,
			OriginalUrl:         req.MasterEntity.Product.OriginalUrl,
			IsActive:            req.MasterEntity.Product.IsActive,
			IsGlobalUpdateLock:  req.MasterEntity.Product.IsGlobalLock,
			IsProductLock:       req.MasterEntity.Product.IsProductLock,
			IsPriceLock:         req.MasterEntity.Product.IsPriceLock,
			IsSpecificationLock: req.MasterEntity.Product.IsSpecificationLock,
			IsImagesLock:        req.MasterEntity.Product.IsImagesLock,
			IsAttributeLock:     req.MasterEntity.Product.IsAttributeLock,
			Factor:              0,
			Uuid:                uuid.NewString(),
		})

		_ = s.images.Create(&entitys.Images{
			ProductId:   productId,
			Type:        true,
			Dir:         consts.DefaultImageDir,
			Filename:    consts.DefaultImageFile,
			OriginalUrl: "defaultBig",
		})

		_ = s.images.Create(&entitys.Images{
			ProductId:   productId,
			Type:        false,
			Dir:         consts.DefaultImageDir,
			Filename:    consts.DefaultImageFile,
			OriginalUrl: "defaultMin",
		})

	} else {
		product.Article = req.MasterEntity.Product.Article
		product.Name = req.MasterEntity.Product.Name
		product.OriginalUrl = req.MasterEntity.Product.OriginalUrl
		product.IsActive = req.MasterEntity.Product.IsActive
		product.IsGlobalUpdateLock = req.MasterEntity.Product.IsGlobalLock
		product.IsProductLock = req.MasterEntity.Product.IsProductLock
		product.IsPriceLock = req.MasterEntity.Product.IsPriceLock
		product.IsSpecificationLock = req.MasterEntity.Product.IsSpecificationLock
		product.IsImagesLock = req.MasterEntity.Product.IsImagesLock
		product.IsAttributeLock = req.MasterEntity.Product.IsAttributeLock
		productId = s.product.Update(product)
	}

	if productId == 0 {
		return nil, errors.New(500, "REQUEST_HAS_UPDATE", "Ошибка в обновлении или создании")
	}

	for _, el := range req.MasterEntity.Price {
		price := s.price.FindByID(el.Id)
		if price == nil || price.ID == 0 {
			_ = s.price.Create(&entitys.Price{
				Price:     float64(el.Price),
				Count:     el.Count,
				Mid:       uuid.NewString(),
				Uuid:      uuid.NewString(),
				Meter:     el.Meter,
				Width:     el.Width,
				Height:    el.Height,
				ProductId: productId,
				OldPrice:  0,
			})
		} else {
			price.OldPrice = float64(price.Price)
			price.Price = float64(el.Price)
			price.Count = el.Count
			price.Mid = el.Mid
			price.Uuid = el.Uuid
			price.Meter = el.Meter
			price.Width = el.Width
			price.Height = el.Height
			price.ProductId = productId
			_ = s.price.Update(price)
		}
	}

	for _, el := range req.MasterEntity.Specification {
		specification := s.specification.FindByID(el.Id)
		if specification == nil || specification.ID == 0 {
			_ = s.specification.Create(&entitys.Specification{
				Name:      el.Name,
				Value:     el.Value,
				ProductId: productId,
			})
		} else {
			specification.Name = el.Name
			specification.Value = el.Value
			specification.ProductId = productId
			_ = s.specification.Update(specification)
		}
	}

	for _, el := range req.MasterEntity.Attribute {
		attribute := s.attribute.FindByID(el.Id)
		if attribute == nil || attribute.ID == 0 {
			_ = s.attribute.Create(&entitys.Attribute{
				Name:      el.Name,
				Value:     el.Value,
				ProductId: productId,
			})
		} else {
			attribute.Name = el.Name
			attribute.Value = el.Value
			attribute.ProductId = productId
			_ = s.attribute.Update(attribute)
		}
	}

	return &pb.SetCarpetReply{
		Success: true,
	}, nil

}
