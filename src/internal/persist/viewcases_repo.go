package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ViewcasesContext interface {
	Create(data *entitys.Viewcases) string
	Update(data *entitys.Viewcases) string
	FindByUUID(string) *entitys.Viewcases
	FindByUUIDAll(uuid string) *entitys.Viewcases
	FindList(limit int64, point int64) []*entitys.Viewcases
	FindListCount() int64
	Drop(uuid string) bool
}

type viewcasesRepo struct {
	em  *Persist
	log *log.Helper
}

func NewViewcasesRepo(p *Persist, logger log.Logger) ViewcasesContext {
	return &viewcasesRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *viewcasesRepo) Drop(uuid string) bool {
	var data *entitys.Viewcases
	qb := p.em.pdb.Where(entitys.Viewcases{
		Uuid: uuid,
	}).First(&data)

	if data.ID != 0 && !util.IsErrQb(qb) {
		p.em.pdb.Delete(&data)
		return true
	}

	return false
}

func (p *viewcasesRepo) Update(data *entitys.Viewcases) string {
	now := time.Now()
	data.UpdateAt = now
	qb := p.em.pdb.Save(&data)
	if qb.Error != nil {
		return ""
	}
	return data.Uuid
}

func (p *viewcasesRepo) Create(data *entitys.Viewcases) string {
	now := time.Now()
	data.UpdateAt = now
	data.CreateAt = now
	qb := p.em.pdb.Create(&data)
	if qb.Error != nil {
		return ""
	}
	return data.Uuid
}

func (p *viewcasesRepo) FindByUUID(uuid string) *entitys.Viewcases {
	var cases *entitys.Viewcases
	qb := p.em.pdb.Where(entitys.Viewcases{
		Uuid:     uuid,
		IsActive: true,
	}).First(&cases)

	if util.IsErrQb(qb) {
		return nil
	}
	return cases
}

func (p *viewcasesRepo) FindByUUIDAll(uuid string) *entitys.Viewcases {
	var cases *entitys.Viewcases
	qb := p.em.pdb.Where(entitys.Viewcases{
		Uuid: uuid,
	}).First(&cases)

	if util.IsErrQb(qb) {
		return nil
	}
	return cases
}

func (p *viewcasesRepo) FindList(limit int64, offset int64) []*entitys.Viewcases {
	var cases []*entitys.Viewcases
	qb := p.em.pdb.Where(entitys.Viewcases{}).Limit(int(limit)).Offset(int(offset)).Find(&cases).Order("id desc")

	if util.IsErrQb(qb) {
		return nil
	}
	return cases
}

func (p *viewcasesRepo) FindListCount() int64 {
	var count int64
	var cases *[]entitys.Viewcases
	qb := p.em.pdb.Where(entitys.Viewcases{}).Find(&cases).Count(&count)

	if util.IsErrQb(qb) {
		return 0
	}
	return count
}
