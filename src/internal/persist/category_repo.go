package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"

	"github.com/go-kratos/kratos/v2/log"
)

type CategoryContext interface {
	FindIsActive() []*entitys.Category
}

type categoryRepo struct {
	em  *Persist
	log *log.Helper
}

func NewCategoryRepo(p *Persist, logger log.Logger) CategoryContext {
	return &categoryRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *categoryRepo) FindIsActive() []*entitys.Category {
	var category []*entitys.Category
	qb := p.em.pdb.Where(entitys.Category{
		IsActive: true,
	}).Find(&category)

	if util.IsErrQb(qb) {
		return nil
	}
	return category
}
