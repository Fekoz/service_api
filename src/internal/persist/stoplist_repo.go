package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"
	"service_api/src/util/consts"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type StopListContext interface {
	GetSlLimitToday(string, string) bool
	CreateOrUpdateSl(string, string)
	FindList(limit int64, offset int64) []*entitys.StopList
	FindListCount() int64
	DropByID(id int64) bool
}

type stopListRepo struct {
	em  *Persist
	log *log.Helper
}

func NewStopListRepo(p *Persist, logger log.Logger) StopListContext {
	return &stopListRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *stopListRepo) GetSlLimitToday(ip string, mth string) bool {
	var sl *entitys.StopList
	now := time.Now()
	per := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)
	to := time.Date(now.Year(), now.Month(), now.Day()+consts.DayPer, now.Hour(), 0, 0, 0, time.UTC)
	qb := p.em.pdb.
		Where("unlock_at BETWEEN ? AND ?", per, to).
		Where("try >= ?", consts.MaxTry).
		Where("ip = ?", ip).
		Where("name = ?", mth).
		First(&sl)
	if qb.Error != nil {
		return false
	}
	if sl != nil {
		return true
	}
	return false
}

func (p *stopListRepo) CreateOrUpdateSl(ip string, mth string) {
	var sl *entitys.StopList
	now := time.Now()
	per := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)
	to := time.Date(now.Year(), now.Month(), now.Day()+consts.DayPer, now.Hour(), 0, 0, 0, time.UTC)
	qb := p.em.pdb.
		Where("unlock_at BETWEEN ? AND ?", per, to).
		Where("ip = ?", ip).
		Where("name = ?", mth).
		First(&sl)
	if qb.Error == nil && sl != nil {
		// update
		sl.Try += 1
		sl.UpdateAt = now
		sl.UnlockAt = to
		p.em.pdb.Save(sl)
		return
	}
	//create
	sl.Ip = ip
	sl.Try = 1
	sl.Type = 1
	sl.UnlockAt = to
	sl.UpdateAt = now
	sl.CreateAt = now
	sl.Name = mth
	p.em.pdb.Create(sl)
}

func (p *stopListRepo) FindList(limit int64, offset int64) []*entitys.StopList {
	var data []*entitys.StopList
	qb := p.em.pdb.Where(entitys.StopList{}).Limit(int(limit)).Offset(int(offset)).Find(&data).Order("id desc")

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *stopListRepo) FindListCount() int64 {
	var count int64
	var data *[]entitys.StopList
	qb := p.em.pdb.Where(entitys.StopList{}).Find(&data).Count(&count)

	if util.IsErrQb(qb) {
		return 0
	}
	return count
}

func (p *stopListRepo) DropByID(id int64) bool {
	var data *entitys.StopList
	qb := p.em.pdb.Where(entitys.StopList{
		ID: uint(id),
	}).First(&data)

	if data.ID != 0 && !util.IsErrQb(qb) {
		p.em.pdb.Delete(&data)
		return true
	}

	return false
}
