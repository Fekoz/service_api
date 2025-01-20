package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type OrderToCarpetContext interface {
	CreateOrder(*entitys.OrderToCarpet)
	Update(data *entitys.OrderToCarpet) int64
	Create(data *entitys.OrderToCarpet) int64
	FindByUUID(uuid string) *entitys.OrderToCarpet
	FindList(limit int64, offset int64) []*entitys.OrderToCarpet
	FindListCount() int64
	FindClientById(id uint) *entitys.Client
	FindPresentById(id uint) *entitys.Presents
	FindProductById(id int64) *entitys.Product
}

type orderToCarpetRepo struct {
	em  *Persist
	log *log.Helper
}

func NewOrderToCarpetRepo(p *Persist, logger log.Logger) OrderToCarpetContext {
	return &orderToCarpetRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *orderToCarpetRepo) CreateOrder(order *entitys.OrderToCarpet) {
	now := time.Now()
	order.UpdateAt = now
	order.CreateAt = now
	p.em.pdb.Create(&order)
}

func (p *orderToCarpetRepo) Update(data *entitys.OrderToCarpet) int64 {
	now := time.Now()
	data.UpdateAt = now
	qb := p.em.pdb.Save(&data)
	if qb.Error != nil {
		return 0
	}
	return int64(data.ID)
}

func (p *orderToCarpetRepo) Create(data *entitys.OrderToCarpet) int64 {
	now := time.Now()
	data.UpdateAt = now
	data.CreateAt = now
	qb := p.em.pdb.Create(&data)
	if qb.Error != nil {
		return 0
	}
	return int64(data.ID)
}

func (p *orderToCarpetRepo) FindByUUID(uuid string) *entitys.OrderToCarpet {
	var data *entitys.OrderToCarpet
	qb := p.em.pdb.Where(&entitys.OrderToCarpet{
		Uuid: uuid,
	}).First(&data)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *orderToCarpetRepo) FindList(limit int64, offset int64) []*entitys.OrderToCarpet {
	var data []*entitys.OrderToCarpet
	qb := p.em.pdb.Where(entitys.OrderToCarpet{}).Limit(int(limit)).Offset(int(offset)).Find(&data).Order("id desc")

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *orderToCarpetRepo) FindListCount() int64 {
	var count int64
	var data *[]entitys.OrderToCarpet
	qb := p.em.pdb.Where(entitys.OrderToCarpet{}).Find(&data).Count(&count)

	if util.IsErrQb(qb) {
		return 0
	}
	return count
}

func (p *orderToCarpetRepo) FindClientById(id uint) *entitys.Client {
	var data *entitys.Client
	qb := p.em.pdb.Where(&entitys.Client{
		ID: id,
	}).First(&data)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *orderToCarpetRepo) FindPresentById(id uint) *entitys.Presents {
	var data *entitys.Presents
	qb := p.em.pdb.Where(&entitys.Presents{
		ID: id,
	}).First(&data)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *orderToCarpetRepo) FindProductById(id int64) *entitys.Product {
	var data *entitys.Product
	qb := p.em.pdb.Where(&entitys.Product{
		ID: id,
	}).First(&data)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}
