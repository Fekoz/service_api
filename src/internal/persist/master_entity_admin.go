package persist

import (
	"fmt"
	"service_api/src/entitys"
	"service_api/src/util"

	pb "bitbucket.org/klaraeng/package_proto/service_api/admin"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type MasterEntityAdminContext interface {
	OneToId(int64) *pb.MasterEntity
	OneToIdIsProduct(*pb.Product) *pb.MasterEntity
	ManyToIds([]int64) []*pb.MasterEntity
	OneToPriceIdLarge(string) *pb.MasterEntityLarge
	GetCarpetCount() int64
	GetCarpetCountParams(*pb.GetCatalogRequest_ProductFind, []*pb.GetCatalogRequest_SpecificationFind, *pb.GetCatalogRequest_PriceFind) int64
	GetCarpetProductList(int64, int64) []*entitys.Product
	GetCarpetProductListParams(int64, int64, *pb.GetCatalogRequest_ProductFind, []*pb.GetCatalogRequest_SpecificationFind, *pb.GetCatalogRequest_PriceFind) []*entitys.Product
	ArticleToProductId(string) int64
	UuidToProductId(string) int64
	GetProductsByMid(string) []*entitys.Product
	GetProductsByUrl(string) []*entitys.Product
}

type masterEntityAdminRepo struct {
	em  *Persist
	log *log.Helper
}

func NewMasterEntityAdmin(p *Persist, logger log.Logger) MasterEntityAdminContext {
	return &masterEntityAdminRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *masterEntityAdminRepo) findToProduct(id int64) *pb.Product {
	var product entitys.Product
	qb := p.em.pdb.First(&product, id)
	if util.IsErrQb(qb) || product.ID == 0 || len(product.Uuid) < 1 {
		return nil
	}
	return &pb.Product{
		Uuid:                product.Uuid,
		IsActive:            product.IsActive,
		Id:                  product.ID,
		Article:             product.Article,
		Name:                product.Name,
		IsGlobalLock:        product.IsGlobalUpdateLock,
		IsProductLock:       product.IsProductLock,
		IsPriceLock:         product.IsPriceLock,
		IsSpecificationLock: product.IsSpecificationLock,
		IsAttributeLock:     product.IsAttributeLock,
		IsImagesLock:        product.IsImagesLock,
	}
}

func (p *masterEntityAdminRepo) findToPrices(id int64) []*pb.Price {
	var prices []entitys.Price
	qb := p.em.pdb.Where(&entitys.Price{ProductId: id}).Find(&prices)
	if util.IsErrQb(qb) || len(prices) < 1 {
		return nil
	}
	var pricesMap []*pb.Price
	for _, element := range prices {
		pricesMap = append(pricesMap, &pb.Price{
			Id:     int64(element.ID),
			Uuid:   element.Uuid,
			Mid:    element.Mid,
			Price:  int64(element.Price),
			Count:  element.Count,
			Meter:  element.Meter,
			Height: element.Height,
			Width:  element.Width,
		})
	}
	return pricesMap
}
func (p *masterEntityAdminRepo) findToPricesOne(id int64) *pb.Price {
	var price entitys.Price
	qb := p.em.pdb.Where(&entitys.Price{ProductId: id}).First(&price)
	if util.IsErrQb(qb) {
		return nil
	}
	return &pb.Price{
		Id:     int64(price.ID),
		Uuid:   price.Uuid,
		Mid:    price.Mid,
		Price:  int64(price.Price),
		Count:  price.Count,
		Meter:  price.Meter,
		Height: price.Height,
		Width:  price.Width,
	}
}

func (p *masterEntityAdminRepo) findToSpecifications(id int64) []*pb.Specification {
	var specifications []entitys.Specification
	qb := p.em.pdb.Where(&entitys.Specification{ProductId: id}).Find(&specifications)
	if util.IsErrQb(qb) || len(specifications) < 1 {
		return nil
	}
	var specificationsMap []*pb.Specification
	for _, element := range specifications {
		specificationsMap = append(specificationsMap, &pb.Specification{
			Id:    int64(element.ID),
			Name:  element.Name,
			Value: element.Value,
		})
	}
	return specificationsMap
}

func (p *masterEntityAdminRepo) findToImages(id int64) []*pb.Images {
	var images []entitys.Images
	qb := p.em.pdb.Where(&entitys.Images{
		ProductId: id,
		Type:      false,
	}).Find(&images)
	if util.IsErrQb(qb) || len(images) < 1 {
		return nil
	}
	var imagesMap []*pb.Images
	for _, element := range images {
		imagesMap = append(imagesMap, &pb.Images{
			Name: element.Filename,
			Dir:  element.Dir,
			Type: element.Type,
		})
	}
	return imagesMap
}
func (p *masterEntityAdminRepo) findToImagesOne(id int64) *pb.Images {
	var images entitys.Images
	qb := p.em.pdb.Where(&entitys.Images{
		ProductId: id,
		Type:      false,
	}).First(&images)
	if util.IsErrQb(qb) {
		return nil
	}
	return &pb.Images{
		Name: images.Filename,
		Dir:  images.Dir,
		Type: images.Type,
	}
}

func (p *masterEntityAdminRepo) findToAttributes(id int64) []*pb.Attribute {
	var attribute []entitys.Attribute
	qb := p.em.pdb.Where(&entitys.Attribute{ProductId: id}).Find(&attribute)
	if util.IsErrQb(qb) || len(attribute) < 1 {
		return nil
	}
	var attributesMap []*pb.Attribute
	for _, element := range attribute {
		attributesMap = append(attributesMap, &pb.Attribute{
			Id:    int64(element.ID),
			Name:  element.Name,
			Value: element.Value,
		})
	}
	return attributesMap
}

func (p *masterEntityAdminRepo) OneToId(id int64) *pb.MasterEntity {
	return &pb.MasterEntity{
		Product:       p.findToProduct(id),
		Price:         p.findToPrices(id),
		Specification: p.findToSpecifications(id),
		Images:        p.findToImages(id),
		Attribute:     p.findToAttributes(id),
	}
}
func (p *masterEntityAdminRepo) OneToIdIsProduct(product *pb.Product) *pb.MasterEntity {
	return &pb.MasterEntity{
		Product:       product,
		Price:         p.findToPrices(product.Id),
		Specification: p.findToSpecifications(product.Id),
		Images:        p.findToImages(product.Id),
		Attribute:     p.findToAttributes(product.Id),
	}
}

func (p *masterEntityAdminRepo) ManyToIds(ids []int64) []*pb.MasterEntity {
	var me []*pb.MasterEntity
	for _, element := range ids {
		me = append(me, &pb.MasterEntity{
			Product:       p.findToProduct(element),
			Price:         p.findToPrices(element),
			Specification: p.findToSpecifications(element),
			Images:        p.findToImages(element),
		})
	}
	if len(me) <= 0 {
		return nil
	}
	return me
}

func (p *masterEntityAdminRepo) OneToPriceIdLarge(uuid string) *pb.MasterEntityLarge {
	var price entitys.Price
	qb := p.em.pdb.Where(entitys.Price{
		Uuid: uuid,
	}).First(&price)
	if util.IsErrQb(qb) || price.ID == 0 {
		return nil
	}
	id := price.ProductId
	return &pb.MasterEntityLarge{
		Product: p.findToProduct(id),
		Price: &pb.Price{
			Id:     int64(price.ID),
			Uuid:   uuid,
			Price:  int64(price.Price),
			Count:  price.Count,
			Meter:  price.Meter,
			Height: price.Height,
			Width:  price.Width,
		},
		Specification: p.findToSpecifications(id),
		Images:        p.findToImagesOne(id),
		Attribute:     p.findToAttributes(id),
	}
}

/**
 * GetCarpet
 */
func (p *masterEntityAdminRepo) GetCarpetCount() int64 {
	var products []entitys.Product
	var count int64
	qb := p.em.pdb.Find(&products).Count(&count)
	if qb.Error != nil {
		return 0
	}
	return count
}
func (p *masterEntityAdminRepo) GetCarpetProductList(limit int64, offset int64) []*entitys.Product {
	var products []*entitys.Product
	qb := p.em.pdb.Order("id").Limit(int(limit)).Offset(int(offset)).Find(&products)
	if qb.Error != nil {
		return nil
	}
	return products
}
func (p *masterEntityAdminRepo) GetCarpetCountParams(
	product *pb.GetCatalogRequest_ProductFind,
	specification []*pb.GetCatalogRequest_SpecificationFind,
	price *pb.GetCatalogRequest_PriceFind) int64 {
	var count int64
	qb := p.em.pdb.Model(&entitys.Product{})
	qb = p.dynamicCarpetJoinParams(qb, product, specification, price)
	qb = qb.Distinct("product.id").Count(&count)
	if qb.Error != nil {
		return 0
	}
	return count
}
func (p *masterEntityAdminRepo) GetCarpetProductListParams(
	limit int64,
	offset int64,
	product *pb.GetCatalogRequest_ProductFind,
	specification []*pb.GetCatalogRequest_SpecificationFind,
	price *pb.GetCatalogRequest_PriceFind) []*entitys.Product {
	var products []*entitys.Product
	qb := p.em.pdb
	qb = p.dynamicCarpetJoinParams(qb, product, specification, price)
	qb = qb.Distinct("product.id").Order("product.id desc").Limit(int(limit)).Offset(int(offset)).Select(
		"product.article",
		"product.id",
		"product.name",
		"product.uuid",
		"product.is_active",
		"product.is_product_lock",
		"product.is_price_lock",
		"product.is_specification_lock",
		"product.is_images_lock",
		"product.is_attribute_lock",
		"product.is_global_update_lock",
	).Find(&products)
	if qb.Error != nil {
		return nil
	}
	return products
}

/**
 * GetCarpet dynamic Param
 */
func (p *masterEntityAdminRepo) dynamicCarpetJoinParams(
	qb *gorm.DB,
	product *pb.GetCatalogRequest_ProductFind,
	specification []*pb.GetCatalogRequest_SpecificationFind,
	price *pb.GetCatalogRequest_PriceFind) *gorm.DB {
	// is specification join and find
	if len(specification) != 0 {
		for index, field := range specification {
			qb = qb.Joins(
				fmt.Sprintf("JOIN specification specification%d ON specification%d.product_id = product.id AND specification%d.name = ? AND specification%d.value = ?",
					index, index, index, index),
				field.Name,
				field.Value)
		}
	}
	// is price join and find
	if price.PriceMax != 0 || len(price.Width) > 0 || len(price.Height) > 0 {
		qb = qb.Joins("JOIN price ON price.product_id = product.id")
		if price.PriceMax != 0 {
			qb = qb.Where("price.price BETWEEN ? AND ?", price.PriceMin, price.PriceMax)
		}
		if len(price.Width) > 0 {
			qb = qb.Where("price.width = ?", price.Width)
		}
		if len(price.Height) > 0 {
			qb = qb.Where("price.height = ?", price.Height)
		}
	}
	// is product find
	if len(product.Name) > 0 {
		qb = qb.Where("product.name LIKE ?", "%"+product.Name+"%")
	}
	if len(product.Article) > 0 {
		qb = qb.Where("product.article LIKE ?", "%"+product.Article+"%")
	}
	return qb
}

/**
 * @param dynamic Article
 */
func (p *masterEntityAdminRepo) ArticleToProductId(article string) int64 {
	var product entitys.Product
	qb := p.em.pdb.Where(entitys.Product{
		Article: article,
	}).First(&product)
	if qb.Error != nil {
		return 0
	}
	return product.ID
}

/**
 * @param dynamic Uuid
 */
func (p *masterEntityAdminRepo) UuidToProductId(uuid string) int64 {
	var product entitys.Product
	qb := p.em.pdb.Where(entitys.Product{
		Uuid: uuid,
	}).First(&product)
	if qb.Error != nil {
		return 0
	}
	return product.ID
}

/**
 * @other serach
 */
func (p *masterEntityAdminRepo) GetProductsByMid(mid string) []*entitys.Product {
	var products []*entitys.Product
	qb := p.em.pdb

	// Соединяем таблицы product и price, и ищем по mid, используем LEFT JOIN
	qb = qb.Joins("LEFT JOIN price pr ON pr.product_id = product.id").
		Where("pr.mid = ?", mid).
		Select(
			"product.article",
			"product.id",
			"product.name",
			"product.uuid",
			"product.is_active",
			"product.is_product_lock",
			"product.is_price_lock",
			"product.is_specification_lock",
			"product.is_images_lock",
			"product.is_attribute_lock",
			"product.is_global_update_lock",
		).Find(&products)

	if qb.Error != nil {
		return nil
	}

	return products
}

func (p *masterEntityAdminRepo) GetProductsByUrl(url string) []*entitys.Product {
	var products []*entitys.Product
	qb := p.em.pdb

	// Поиск продуктов по url
	qb = qb.Where("product.original_url = ?", url).
		Select(
			"product.article",
			"product.id",
			"product.name",
			"product.uuid",
			"product.is_active",
			"product.is_product_lock",
			"product.is_price_lock",
			"product.is_specification_lock",
			"product.is_images_lock",
			"product.is_attribute_lock",
			"product.is_global_update_lock",
		).Find(&products)

	if qb.Error != nil {
		return nil
	}
	return products
}
