package persist

import (
	"service_api/src/entitys"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type SenderContext interface {
	CreateFormLanding(uuid string, clientId uint)
	CreatePageItemCarpet(uuid string, clientId uint)
	CreatePageItemCarpetPrice(uuid string, clientId uint)
}

type senderRepo struct {
	em  *Persist
	log *log.Helper
}

func NewSenderRepo(p *Persist, logger log.Logger) SenderContext {
	return &senderRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *senderRepo) CreateFormLanding(uuid string, clientId uint) {
	var sender *entitys.Sender
	now := time.Now()
	sender = &entitys.Sender{
		Uuid:           uuid,
		ClientId:       clientId,
		Type:           1, // Тип - Подбор ковра
		IsSend:         false,
		MailTemplateId: 1, // Тип сообщения - Подбор ковра
		UpdateAt:       now,
		CreateAt:       now,
	}
	p.em.pdb.Create(&sender)
}

func (p *senderRepo) CreatePageItemCarpet(uuid string, clientId uint) {
	var sender *entitys.Sender
	now := time.Now()
	sender = &entitys.Sender{
		Uuid:           uuid,
		ClientId:       clientId,
		Type:           2, // Тип - Заказ по товару
		IsSend:         false,
		MailTemplateId: 2, // Тип сообщения - Заказ по товару
		UpdateAt:       now,
		CreateAt:       now,
	}
	p.em.pdb.Create(&sender)
}

func (p *senderRepo) CreatePageItemCarpetPrice(uuid string, clientId uint) {
	var sender *entitys.Sender
	now := time.Now()
	sender = &entitys.Sender{
		Uuid:           uuid,
		ClientId:       clientId,
		Type:           3, // Тип - Заказ по продукту с размерами
		IsSend:         false,
		MailTemplateId: 3, // Тип сообщения - Заказ по продукту с размерами
		UpdateAt:       now,
		CreateAt:       now,
	}
	p.em.pdb.Create(&sender)
}
