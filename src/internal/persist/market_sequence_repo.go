package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type MarketSequenceContext interface {
	SetActiveByMid(mid string, isActive bool, isCounter bool, counterPkg int64)
	FindByMid(mid string) bool
	FindByActive(bool) []*entitys.MarketSequence
	FindList(limit int64, point int64) []*entitys.MarketSequence
	FindListCount() int64
}

type marketSequenceRepo struct {
	em  *Persist
	log *log.Helper
}

func NewMarketSequenceRepo(p *Persist, logger log.Logger) MarketSequenceContext {
	return &marketSequenceRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *marketSequenceRepo) SetActiveByMid(mid string, isActive bool, isCounter bool, counterPkg int64) {
	now := time.Now()
	var data *entitys.MarketSequence
	qb := p.em.pdb.Where(entitys.MarketSequence{
		Mid: mid,
	}).First(&data)

	if !util.IsErrQb(qb) && data.ID != 0 {
		data.IsActive = isActive
		data.IsCounter = isCounter
		data.CounterPkg = counterPkg
		data.UpdateAt = now
		p.em.pdb.Save(&data)
		return
	}

	data.IsActive = isActive
	data.Mid = mid
	data.Name = "market.locker"
	data.IsCounter = isCounter
	data.CounterPkg = counterPkg
	data.CreateAt = now
	data.UpdateAt = now
	p.em.pdb.Create(&data)
}

func (p *marketSequenceRepo) FindByMid(mid string) bool {
	var data *entitys.MarketSequence
	p.em.pdb.Where(entitys.MarketSequence{
		Mid: mid,
	}).First(&data)

	if len(data.Mid) > 0 {
		return true
	}

	return false
}

func (p *marketSequenceRepo) FindByActive(isActive bool) []*entitys.MarketSequence {
	var data []*entitys.MarketSequence
	qb := p.em.pdb.Where(entitys.MarketSequence{
		IsActive: isActive,
	}).Find(&data)

	if util.IsErrQb(qb) {
		return data
	}
	return data
}

func (p *marketSequenceRepo) FindList(limit int64, offset int64) []*entitys.MarketSequence {
	var data []*entitys.MarketSequence
	qb := p.em.pdb.Where(entitys.MarketSequence{}).Limit(int(limit)).Offset(int(offset)).Find(&data).Order("id desc")

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *marketSequenceRepo) FindListCount() int64 {
	var count int64
	var data *[]entitys.MarketSequence
	qb := p.em.pdb.Where(entitys.MarketSequence{}).Find(&data).Count(&count)

	if util.IsErrQb(qb) {
		return 0
	}
	return count
}
