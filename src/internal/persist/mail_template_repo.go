package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type MailTemplateContext interface {
	SetByID(id int64, title string, message string) bool
	FindByID(id int64) *entitys.MailTemplate
	FindList(limit int64, point int64) []*entitys.MailTemplate
	FindListCount() int64
}

type mailTemplateRepo struct {
	em  *Persist
	log *log.Helper
}

func NewMailTemplateRepo(p *Persist, logger log.Logger) MailTemplateContext {
	return &mailTemplateRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *mailTemplateRepo) SetByID(id int64, title string, message string) bool {
	now := time.Now()
	var data *entitys.MailTemplate
	qb := p.em.pdb.Where(entitys.MailTemplate{
		ID: uint(id),
	}).First(&data)

	if data.ID != 0 && !util.IsErrQb(qb) {
		data.Title = title
		data.Message = message
		data.UpdateAt = now
		p.em.pdb.Save(&data)
		return true
	}

	return false
}

func (p *mailTemplateRepo) FindByID(id int64) *entitys.MailTemplate {
	var data *entitys.MailTemplate
	qb := p.em.pdb.Where(entitys.MailTemplate{
		ID: uint(id),
	}).First(&data)

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *mailTemplateRepo) FindList(limit int64, offset int64) []*entitys.MailTemplate {
	var data []*entitys.MailTemplate
	qb := p.em.pdb.Where(entitys.MailTemplate{}).Limit(int(limit)).Offset(int(offset)).Find(&data).Order("id desc")

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *mailTemplateRepo) FindListCount() int64 {
	var count int64
	var data *[]entitys.MailTemplate
	qb := p.em.pdb.Where(entitys.MailTemplate{}).Find(&data).Count(&count)

	if util.IsErrQb(qb) {
		return 0
	}
	return count
}
