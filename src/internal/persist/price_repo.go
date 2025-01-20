package persist

import (
	"fmt"
	"service_api/src/dtos"
	"service_api/src/entitys"
	"service_api/src/util"

	pa "bitbucket.org/klaraeng/package_proto/service_api/admin"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/log"
)

type PriceContext interface {
	Update(data *entitys.Price) bool
	Create(data *entitys.Price) bool
	FindByID(id int64) *entitys.Price
	FindByUUID(uuid string) *entitys.Price
	FindByMID(mid string) *entitys.Price
	RangePriceList(int64) *pb.MasterEntityCrop_PriceRange
	RangePriceListAdmin(int64) *pa.MasterEntityCrop_PriceRange
	OneToProductIdPrice(int64) []entitys.Price
	FindWithPrice([]*dtos.CatalogFields) []entitys.Price
}

type priceRepo struct {
	em  *Persist
	log *log.Helper
}

func NewPriceRepo(p *Persist, logger log.Logger) PriceContext {
	return &priceRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *priceRepo) Update(data *entitys.Price) bool {
	qb := p.em.pdb.Save(&data)
	if qb.Error != nil {
		return false
	}
	return true
}

func (p *priceRepo) Create(data *entitys.Price) bool {
	qb := p.em.pdb.Create(&data)
	if qb.Error != nil {
		return false
	}
	return true
}

func (p *priceRepo) FindByID(id int64) *entitys.Price {
	var data *entitys.Price
	qb := p.em.pdb.First(&data, id)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *priceRepo) FindByUUID(uuid string) *entitys.Price {
	var data *entitys.Price
	qb := p.em.pdb.Where(&entitys.Price{Uuid: uuid}).Find(&data)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *priceRepo) FindByMID(mid string) *entitys.Price {
	var data *entitys.Price
	qb := p.em.pdb.Where(&entitys.Price{Mid: mid}).Find(&data)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *priceRepo) RangePriceList(productId int64) *pb.MasterEntityCrop_PriceRange {
	var prices []*entitys.Price
	qb := p.em.pdb.Where(&entitys.Images{ProductId: productId}).Find(&prices)
	if util.IsErrQb(qb) {
		return nil
	}
	var priceStart int64
	var priceEnd int64
	var count int64
	for index, el := range prices {
		if index == 0 {
			priceStart = int64(el.Price)
			priceEnd = int64(el.Price)
		}
		if int64(el.Price) < priceStart {
			priceStart = int64(el.Price)
		}
		if int64(el.Price) > priceEnd {
			priceEnd = int64(el.Price)
		}
		count += el.Count
	}
	return &pb.MasterEntityCrop_PriceRange{
		PriceStart: priceStart,
		PriceEnd:   priceEnd,
		Count:      count,
	}
}

func (p *priceRepo) RangePriceListAdmin(productId int64) *pa.MasterEntityCrop_PriceRange {
	var prices []*entitys.Price
	qb := p.em.pdb.Where(&entitys.Images{ProductId: productId}).Find(&prices)
	if util.IsErrQb(qb) {
		return nil
	}
	var priceStart int64
	var priceEnd int64
	var count int64
	for index, el := range prices {
		if index == 0 {
			priceStart = int64(el.Price)
			priceEnd = int64(el.Price)
		}
		if int64(el.Price) < priceStart {
			priceStart = int64(el.Price)
		}
		if int64(el.Price) > priceEnd {
			priceEnd = int64(el.Price)
		}
		count += el.Count
	}
	return &pa.MasterEntityCrop_PriceRange{
		PriceStart: priceStart,
		PriceEnd:   priceEnd,
		Count:      count,
	}
}

func (p *priceRepo) OneToProductIdPrice(id int64) []entitys.Price {
	var prices []entitys.Price
	qb := p.em.pdb.Where(&entitys.Price{
		ProductId: id,
	}).Find(&prices)
	if util.IsErrQb(qb) || len(prices) == 0 {
		return nil
	}
	return prices
}

func (p *priceRepo) FindWithPrice(fields []*dtos.CatalogFields) []entitys.Price {
	var prices []entitys.Price
	qb := p.em.pdb
	for _, field := range fields {
		qb = qb.Where(fmt.Sprintf("%s %s ?", field.Name, field.Attribute), field.Value)
	}
	qb = qb.Order("id desc").Find(&prices)
	if util.IsErrQb(qb) || len(prices) == 0 {
		return nil
	}
	return prices
}
