package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"
	"service_api/src/util/consts"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type OrderContext interface {
	FindListCount() int64
	FindList(limit int64, offset int64) []*entitys.Order
	Update(data *entitys.Order) int64
	Create(data *entitys.Order) int64
	FindByUUID(uuid string) *entitys.Order
	FindOrderFieldById(id uint) []*entitys.OrderField
	GetPresentWithCode(string) (uint, error)
	CreateOrder(order *entitys.Order) (uint, error)
	CreateOrderField(string, string, uint)
	CreateOrderFields(fields []entitys.OrderField)
	FindClientById(id uint) *entitys.Client
	FindPresentById(id uint) *entitys.Presents
}

type orderRepo struct {
	em  *Persist
	log *log.Helper
}

func NewOrderRepo(p *Persist, logger log.Logger) OrderContext {
	return &orderRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *orderRepo) Update(data *entitys.Order) int64 {
	now := time.Now()
	data.UpdateAt = now
	qb := p.em.pdb.Save(&data)
	if qb.Error != nil {
		return 0
	}
	return int64(data.ID)
}

func (p *orderRepo) Create(data *entitys.Order) int64 {
	now := time.Now()
	data.UpdateAt = now
	data.CreateAt = now
	qb := p.em.pdb.Create(&data)
	if qb.Error != nil {
		return 0
	}
	return int64(data.ID)
}

func (p *orderRepo) FindByUUID(uuid string) *entitys.Order {
	var data *entitys.Order
	qb := p.em.pdb.Where(&entitys.Order{
		Uuid: uuid,
	}).First(&data)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *orderRepo) FindList(limit int64, offset int64) []*entitys.Order {
	var data []*entitys.Order
	qb := p.em.pdb.Where(entitys.Order{}).Limit(int(limit)).Offset(int(offset)).Find(&data).Order("id desc")

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *orderRepo) FindListCount() int64 {
	var count int64
	var data *[]entitys.Order
	qb := p.em.pdb.Where(entitys.Order{}).Find(&data).Count(&count)

	if util.IsErrQb(qb) {
		return 0
	}
	return count
}

func (p *orderRepo) FindOrderFieldById(id uint) []*entitys.OrderField {
	var data []*entitys.OrderField
	qb := p.em.pdb.Where(entitys.OrderField{
		OrderID: id,
	}).Find(&data)

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *orderRepo) GetPresentWithCode(code string) (uint, error) {
	var presents *entitys.Presents
	qb := p.em.pdb.Where(&entitys.Presents{
		Key: code,
	}).First(&presents)

	if qb.Error != nil {
		return 0, qb.Error
	}
	return presents.ID, nil
}

func (p *orderRepo) CreateOrder(order *entitys.Order) (uint, error) {
	now := time.Now()
	order.UpdateAt = now
	order.CreateAt = now
	qb := p.em.pdb.Create(&order)
	if qb.Error != nil {
		return 0, qb.Error
	}
	return order.ID, nil
}

func (p *orderRepo) CreateOrderField(key string, val string, id uint) {
	em := entitys.OrderField{
		OrderID: id,
		Key:     key,
		Value:   val,
	}
	p.em.pdb.Create(&em)
}

func (p *orderRepo) CreateOrderFields(fields []entitys.OrderField) {
	p.em.pdb.CreateInBatches(&fields, consts.PacketSizeInsertOrderField)
}

func (p *orderRepo) FindClientById(id uint) *entitys.Client {
	var data *entitys.Client
	qb := p.em.pdb.Where(&entitys.Client{
		ID: id,
	}).First(&data)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *orderRepo) FindPresentById(id uint) *entitys.Presents {
	var data *entitys.Presents
	qb := p.em.pdb.Where(&entitys.Presents{
		ID: id,
	}).First(&data)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}
