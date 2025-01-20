package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ShowcasesContext interface {
	Drop(uuid string) bool
	Create(data *entitys.Showcases) string
	Update(data *entitys.Showcases) string
	FindByUUID(string) *entitys.Showcases
	FindByUUIDAll(uuid string) *entitys.Showcases
	FindList(limit int64, point int64) []*entitys.Showcases
	FindListCount() int64
}

type showcasesRepo struct {
	em  *Persist
	log *log.Helper
}

func NewShowcasesRepo(p *Persist, logger log.Logger) ShowcasesContext {
	return &showcasesRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *showcasesRepo) Drop(uuid string) bool {
	var data *entitys.Showcases
	qb := p.em.pdb.Where(entitys.Showcases{
		Uuid: uuid,
	}).First(&data)

	if data.ID != 0 && !util.IsErrQb(qb) {
		p.em.pdb.Delete(&data)
		return true
	}

	return false
}

func (p *showcasesRepo) Update(data *entitys.Showcases) string {
	now := time.Now()
	data.UpdateAt = now
	qb := p.em.pdb.Save(&data)
	if qb.Error != nil {
		return ""
	}
	return data.Uuid
}

func (p *showcasesRepo) Create(data *entitys.Showcases) string {
	now := time.Now()
	data.UpdateAt = now
	data.CreateAt = now
	qb := p.em.pdb.Create(&data)
	if qb.Error != nil {
		return ""
	}
	return data.Uuid
}

func (p *showcasesRepo) FindByUUID(uuid string) *entitys.Showcases {
	var showcases *entitys.Showcases
	qb := p.em.pdb.Where(entitys.Showcases{
		Uuid:     uuid,
		IsActive: true,
	}).First(&showcases)

	if util.IsErrQb(qb) {
		return nil
	}
	return showcases
}

func (p *showcasesRepo) FindByUUIDAll(uuid string) *entitys.Showcases {
	var showcases *entitys.Showcases
	qb := p.em.pdb.Where(entitys.Showcases{
		Uuid: uuid,
	}).First(&showcases)

	if util.IsErrQb(qb) {
		return nil
	}
	return showcases
}

func (p *showcasesRepo) FindList(limit int64, offset int64) []*entitys.Showcases {
	var showcases []*entitys.Showcases
	qb := p.em.pdb.Where(entitys.Showcases{}).Limit(int(limit)).Offset(int(offset)).Find(&showcases).Order("id desc")

	if util.IsErrQb(qb) {
		return nil
	}
	return showcases
}

func (p *showcasesRepo) FindListCount() int64 {
	var count int64
	var showcases *[]entitys.Showcases
	qb := p.em.pdb.Where(entitys.Showcases{}).Find(&showcases).Count(&count)

	if util.IsErrQb(qb) {
		return 0
	}
	return count
}
