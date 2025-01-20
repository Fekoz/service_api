package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type OptionsContext interface {
	Update(*entitys.Options) bool
	Find(string) *entitys.Options
	List() []*entitys.Options
}

type optionsRepo struct {
	em  *Persist
	log *log.Helper
}

func NewOptionsRepo(p *Persist, logger log.Logger) OptionsContext {
	return &optionsRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *optionsRepo) Update(data *entitys.Options) bool {
	now := time.Now()
	data.UpdateAt = now
	qb := p.em.pdb.Save(&data)
	if qb.Error != nil {
		return false
	}
	return true
}

func (p *optionsRepo) Find(name string) *entitys.Options {
	var data *entitys.Options
	qb := p.em.pdb.Where(&entitys.Options{
		Name: name,
	}).First(&data)

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *optionsRepo) List() []*entitys.Options {
	var data []*entitys.Options
	qb := p.em.pdb.Where(entitys.Options{}).Find(&data).Order("id desc")

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}
