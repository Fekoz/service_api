package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type PriceExportDynamicContext interface {
	Update(data *entitys.PriceExportDynamic)
	Create(minWidth int64, maxWidth int64, minHeight int64, maxHeight int64, pckDepth int64, pckWidth int64, pckHeight int64) int64
	Drop(data *entitys.PriceExportDynamic)
	FindList(limit int64, offset int64) []*entitys.PriceExportDynamic
	FindListCount() int64
	FindByID(id int64) *entitys.PriceExportDynamic
}

type priceExportDynamicRepo struct {
	em  *Persist
	log *log.Helper
}

func NewPriceExportDynamicRepo(p *Persist, logger log.Logger) PriceExportDynamicContext {
	return &priceExportDynamicRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *priceExportDynamicRepo) Update(data *entitys.PriceExportDynamic) {
	now := time.Now()
	data.UpdateAt = now
	p.em.pdb.Save(&data)
}

func (p *priceExportDynamicRepo) Create(
	minWidth int64,
	maxWidth int64,
	minHeight int64,
	maxHeight int64,
	pckDepth int64,
	pckWidth int64,
	pckHeight int64,
) int64 {
	entity := &entitys.PriceExportDynamic{
		MinWidth:      minWidth,
		MaxWidth:      maxWidth,
		MinHeight:     minHeight,
		MaxHeight:     maxHeight,
		PackageDepth:  pckDepth,
		PackageWidth:  pckWidth,
		PackageHeight: pckHeight,
	}

	now := time.Now()
	entity.UpdateAt = now
	entity.CreateAt = now
	qb := p.em.pdb.Create(&entity)
	if qb.Error != nil {
		return 0
	}
	return int64(entity.ID)
}

func (p *priceExportDynamicRepo) Drop(data *entitys.PriceExportDynamic) {
	p.em.pdb.Delete(&data)
}

func (p *priceExportDynamicRepo) FindList(limit int64, offset int64) []*entitys.PriceExportDynamic {
	var data []*entitys.PriceExportDynamic
	qb := p.em.pdb.Where(entitys.PriceExportDynamic{}).Limit(int(limit)).Offset(int(offset)).Find(&data).Order("id desc")

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *priceExportDynamicRepo) FindListCount() int64 {
	var count int64
	var data *[]entitys.PriceExportDynamic
	qb := p.em.pdb.Where(entitys.PriceExportDynamic{}).Find(&data).Count(&count)

	if util.IsErrQb(qb) {
		return 0
	}
	return count
}

func (p *priceExportDynamicRepo) FindByID(id int64) *entitys.PriceExportDynamic {
	var data *entitys.PriceExportDynamic
	qb := p.em.pdb.First(&data, id)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}
