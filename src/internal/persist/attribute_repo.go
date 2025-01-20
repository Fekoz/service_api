package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"

	"github.com/go-kratos/kratos/v2/log"
)

type AttributeContext interface {
	Update(data *entitys.Attribute) bool
	Create(data *entitys.Attribute) bool
	FindByID(id int64) *entitys.Attribute
	OneToProductIdAttribute(id int64) []entitys.Attribute
}

type attributeRepo struct {
	em  *Persist
	log *log.Helper
}

func NewAttributeRepo(p *Persist, logger log.Logger) AttributeContext {
	return &attributeRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *attributeRepo) Update(data *entitys.Attribute) bool {
	qb := p.em.pdb.Save(&data)
	if qb.Error != nil {
		return false
	}
	return true
}

func (p *attributeRepo) Create(data *entitys.Attribute) bool {
	qb := p.em.pdb.Create(&data)
	if qb.Error != nil {
		return false
	}
	return true
}

func (p *attributeRepo) FindByID(id int64) *entitys.Attribute {
	var data *entitys.Attribute
	qb := p.em.pdb.First(&data, id)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *attributeRepo) OneToProductIdAttribute(id int64) []entitys.Attribute {
	var attribute []entitys.Attribute
	qb := p.em.pdb.Where(&entitys.Attribute{
		ProductId: id,
	}).Find(&attribute)
	if util.IsErrQb(qb) || len(attribute) == 0 {
		return nil
	}
	return attribute
}
