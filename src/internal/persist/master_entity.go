package persist

import (
	"fmt"
	"service_api/src/entitys"
	"service_api/src/util"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type MasterEntityContext interface {
	OneToId(int64) *pb.MasterEntity
	OneToIdIsProduct(*pb.Product) *pb.MasterEntity
	ManyToIds([]int64) []*pb.MasterEntity
	OneToPriceIdLarge(string) *pb.MasterEntityLarge
	GetCarpetCount() int64
	GetCarpetCountParams(*pb.GetCatalogRequest_ProductFind, []*pb.GetCatalogRequest_SpecificationFind, *pb.GetCatalogRequest_PriceFind, *pb.GetCatalogRequest_SizeFind) int64
	GetCarpetProductList(int64, int64) []*entitys.Product
	GetCarpetProductListParams(int64, int64, *pb.GetCatalogRequest_ProductFind, []*pb.GetCatalogRequest_SpecificationFind, *pb.GetCatalogRequest_PriceFind, *pb.GetCatalogRequest_SizeFind) []*entitys.Product
	ArticleToProductId(string) int64
	UuidToProductId(string) int64
	ProductIdToArticle(id int64) string
}

type masterEntityRepo struct {
	em  *Persist
	log *log.Helper
}

func NewMasterEntity(p *Persist, logger log.Logger) MasterEntityContext {
	return &masterEntityRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *masterEntityRepo) findToProduct(id int64) *pb.Product {
	var product entitys.Product
	qb := p.em.pdb.First(&product, id)
	if util.IsErrQb(qb) || product.ID == 0 || len(product.Uuid) < 1 {
		return nil
	}
	return &pb.Product{
		Uuid:     product.Uuid,
		IsActive: product.IsActive,
		Id:       product.ID,
		Article:  product.Article,
		Name:     product.Name,
	}
}

func (p *masterEntityRepo) findToPrices(id int64) []*pb.Price {
	var prices []entitys.Price
	qb := p.em.pdb.Where(&entitys.Price{ProductId: id}).Find(&prices)
	if util.IsErrQb(qb) || len(prices) < 1 {
		return nil
	}
	var pricesMap []*pb.Price
	for _, element := range prices {
		pricesMap = append(pricesMap, &pb.Price{
			Id:     int64(element.ID),
			Price:  int64(element.Price),
			Count:  element.Count,
			Meter:  element.Meter,
			Height: element.Height,
			Width:  element.Width,
			Uuid:   element.Uuid,
			Mid:    element.Mid,
		})
	}
	return pricesMap
}
func (p *masterEntityRepo) findToPricesOne(id int64) *pb.Price {
	var price entitys.Price
	qb := p.em.pdb.Where(&entitys.Price{ProductId: id}).First(&price)
	if util.IsErrQb(qb) {
		return nil
	}
	return &pb.Price{
		Id:     int64(price.ID),
		Price:  int64(price.Price),
		Count:  price.Count,
		Meter:  price.Meter,
		Height: price.Height,
		Width:  price.Width,
	}
}

func (p *masterEntityRepo) findToSpecifications(id int64) []*pb.Specification {
	var specifications []entitys.Specification
	qb := p.em.pdb.Where(&entitys.Specification{ProductId: id}).Find(&specifications)
	if util.IsErrQb(qb) || len(specifications) < 1 {
		return nil
	}
	var specificationsMap []*pb.Specification
	for _, element := range specifications {
		specificationsMap = append(specificationsMap, &pb.Specification{
			Name:  element.Name,
			Value: element.Value,
		})
	}
	return specificationsMap
}

func (p *masterEntityRepo) findToImages(id int64) []*pb.Images {
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
func (p *masterEntityRepo) findToImagesOne(id int64) *pb.Images {
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

func (p *masterEntityRepo) findToAttributes(id int64) []*pb.Attribute {
	var attribute []entitys.Attribute
	qb := p.em.pdb.Where(&entitys.Attribute{ProductId: id}).Find(&attribute)
	if util.IsErrQb(qb) || len(attribute) < 1 {
		return nil
	}
	var attributesMap []*pb.Attribute
	for _, element := range attribute {
		attributesMap = append(attributesMap, &pb.Attribute{
			Name:  element.Name,
			Value: element.Value,
		})
	}
	return attributesMap
}

func (p *masterEntityRepo) OneToId(id int64) *pb.MasterEntity {
	return &pb.MasterEntity{
		Product:       p.findToProduct(id),
		Price:         p.findToPrices(id),
		Specification: p.findToSpecifications(id),
		Images:        p.findToImages(id),
		Attribute:     p.findToAttributes(id),
	}
}
func (p *masterEntityRepo) OneToIdIsProduct(product *pb.Product) *pb.MasterEntity {
	return &pb.MasterEntity{
		Product:       product,
		Price:         p.findToPrices(product.Id),
		Specification: p.findToSpecifications(product.Id),
		Images:        p.findToImages(product.Id),
		Attribute:     p.findToAttributes(product.Id),
	}
}

func (p *masterEntityRepo) ManyToIds(ids []int64) []*pb.MasterEntity {
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

func (p *masterEntityRepo) OneToPriceIdLarge(uuid string) *pb.MasterEntityLarge {
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
func (p *masterEntityRepo) GetCarpetCount() int64 {
	var products []entitys.Product
	var count int64
	qb := p.em.pdb.Find(&products).Count(&count)
	if qb.Error != nil {
		return 0
	}
	return count
}
func (p *masterEntityRepo) GetCarpetProductList(limit int64, offset int64) []*entitys.Product {
	var products []*entitys.Product
	qb := p.em.pdb.Order("id").Limit(int(limit)).Offset(int(offset)).Find(&products)
	if qb.Error != nil {
		return nil
	}
	return products
}
func (p *masterEntityRepo) GetCarpetCountParams(
	product *pb.GetCatalogRequest_ProductFind,
	specification []*pb.GetCatalogRequest_SpecificationFind,
	price *pb.GetCatalogRequest_PriceFind,
	size *pb.GetCatalogRequest_SizeFind,
) int64 {
	var products []entitys.Product
	var count int64
	qb := p.em.pdb
	qb = p.dynamicCarpetJoinParams(qb, product, specification, price, size)
	qb = qb.Distinct("product.id").Order("product.id desc").Find(&products).Count(&count)
	if qb.Error != nil {
		return 0
	}
	return count
}
func (p *masterEntityRepo) GetCarpetProductListParams(
	limit int64,
	offset int64,
	product *pb.GetCatalogRequest_ProductFind,
	specification []*pb.GetCatalogRequest_SpecificationFind,
	price *pb.GetCatalogRequest_PriceFind,
	size *pb.GetCatalogRequest_SizeFind,
) []*entitys.Product {
	var products []*entitys.Product
	qb := p.em.pdb
	qb = p.dynamicCarpetJoinParams(qb, product, specification, price, size)
	qb = qb.Distinct("product.id").Order("product.update_at asc").Limit(int(limit)).Offset(int(offset)).Select(
		"product.article",
		"product.id",
		"product.name",
		"product.uuid",
		"product.is_active",
		"product.update_at",
	).Find(&products)
	if qb.Error != nil {
		return nil
	}
	return products
}

/**
 * GetCarpet dynamic Param
 */
func (p *masterEntityRepo) dynamicCarpetJoinParams(
	qb *gorm.DB,
	product *pb.GetCatalogRequest_ProductFind,
	specification []*pb.GetCatalogRequest_SpecificationFind,
	price *pb.GetCatalogRequest_PriceFind,
	size *pb.GetCatalogRequest_SizeFind,
) *gorm.DB {
	if product == nil {
		product = &pb.GetCatalogRequest_ProductFind{}
	}
	if specification == nil {
		specification = []*pb.GetCatalogRequest_SpecificationFind{}
	}
	if price == nil {
		price = &pb.GetCatalogRequest_PriceFind{}
	}
	if size == nil {
		size = &pb.GetCatalogRequest_SizeFind{}
	}

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
	if price.PriceMax != 0 || len(size.Width) > 0 || len(size.Height) > 0 {
		qb = qb.Joins("JOIN price ON price.product_id = product.id")
		if price.PriceMax != 0 {
			qb = qb.Where("price.price BETWEEN ? AND ?", price.PriceMin, price.PriceMax)
		}
	}

	// is product find
	if len(product.Name) > 0 {
		qb = qb.Where("product.name LIKE ?", "%"+product.Name+"%")
	}
	if len(product.Article) > 0 {
		qb = qb.Where("product.article LIKE ?", "%"+product.Article+"%")
	}

	// size find
	if len(size.Width) > 0 {
		qb = qb.Where("price.width IN ?", size.Width)
	}
	// size find
	if len(size.Height) > 0 {
		qb = qb.Where("price.height IN ?", size.Height)
	}

	if len(product.Name) > 0 {
		qb = qb.Where("product.name LIKE ?", "%"+product.Name+"%")
	}
	if len(product.Article) > 0 {
		qb = qb.Where("product.article LIKE ?", "%"+product.Article+"%")
	}

	qb = qb.Where("product.is_active = true")
	return qb
}

/**
 * @param dynamic Article
 */
func (p *masterEntityRepo) ArticleToProductId(article string) int64 {
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
func (p *masterEntityRepo) UuidToProductId(uuid string) int64 {
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
 * @param dynamic ProductId
 */
func (p *masterEntityRepo) ProductIdToArticle(id int64) string {
	var product entitys.Product
	qb := p.em.pdb.Where(entitys.Product{
		ID: id,
	}).First(&product)
	if qb.Error != nil {
		return ""
	}
	return product.Article
}
