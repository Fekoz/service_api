package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ProductContext interface {
	Update(data *entitys.Product) int64
	Create(data *entitys.Product) int64
	FindByID(int64) *entitys.Product
	FindByArticle(string) *entitys.Product
	FindByUuid(string) *entitys.Product
	OneToProductIdProduct(int64) *entitys.Product
}

type productRepo struct {
	em  *Persist
	log *log.Helper
}

func NewProductRepo(p *Persist, logger log.Logger) ProductContext {
	return &productRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *productRepo) Update(data *entitys.Product) int64 {
	now := time.Now()
	data.UpdateAt = now
	qb := p.em.pdb.Save(&data)
	if qb.Error != nil {
		return 0
	}
	return data.ID
}

func (p *productRepo) Create(data *entitys.Product) int64 {
	now := time.Now()
	data.UpdateAt = now
	data.CreateAt = now
	qb := p.em.pdb.Create(&data)
	if qb.Error != nil {
		return 0
	}
	return data.ID
}

func (p *productRepo) FindByID(id int64) *entitys.Product {
	var data *entitys.Product
	qb := p.em.pdb.First(&data, id)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *productRepo) FindByArticle(article string) *entitys.Product {
	var product *entitys.Product
	qb := p.em.pdb.Where(entitys.Product{
		Article: article,
	}).First(&product)
	if util.IsErrQb(qb) {
		return nil
	}
	return product
}

func (p *productRepo) FindByUuid(uuid string) *entitys.Product {
	var product *entitys.Product
	qb := p.em.pdb.Where(entitys.Product{
		Uuid: uuid,
	}).First(&product)
	if util.IsErrQb(qb) {
		return nil
	}
	return product
}

func (p *productRepo) OneToProductIdProduct(id int64) *entitys.Product {
	var product *entitys.Product
	qb := p.em.pdb.Where(entitys.Product{
		ID: id,
	}).Find(&product)
	if util.IsErrQb(qb) {
		return nil
	}
	return product
}
