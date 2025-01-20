package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"

	"github.com/go-kratos/kratos/v2/log"
)

type SpecificationContext interface {
	Update(data *entitys.Specification) bool
	Create(data *entitys.Specification) bool
	FindByID(id int64) *entitys.Specification
	OneToProductIdSpecification(id int64) []entitys.Specification
}

type specificationRepo struct {
	em  *Persist
	log *log.Helper
}

func NewSpecificationRepo(p *Persist, logger log.Logger) SpecificationContext {
	return &specificationRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *specificationRepo) Update(data *entitys.Specification) bool {
	qb := p.em.pdb.Save(&data)
	if qb.Error != nil {
		return false
	}
	return true
}

func (p *specificationRepo) Create(data *entitys.Specification) bool {
	qb := p.em.pdb.Create(&data)
	if qb.Error != nil {
		return false
	}
	return true
}

func (p *specificationRepo) FindByID(id int64) *entitys.Specification {
	var data *entitys.Specification
	qb := p.em.pdb.First(&data, id)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *specificationRepo) OneToProductIdSpecification(id int64) []entitys.Specification {
	var specification []entitys.Specification
	qb := p.em.pdb.Where(&entitys.Specification{
		ProductId: id,
	}).Find(&specification)
	if util.IsErrQb(qb) || len(specification) == 0 {
		return nil
	}
	return specification
}
