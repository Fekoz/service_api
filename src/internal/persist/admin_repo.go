package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminContext interface {
	IsAdmin(string, string, string) bool
	GetAdmin(string, string, string) *entitys.Admin
	FindAdmin(uint) *entitys.Admin
	FindList(limit int64, offset int64) []*entitys.Admin
	FindListCount() int64
	SetByID(login string, name string, email string, pass string, isActive bool, role int64, protection int64) bool
}

type adminRepo struct {
	em  *Persist
	log *log.Helper
}

func NewAdminRepo(p *Persist, logger log.Logger) AdminContext {
	return &adminRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *adminRepo) IsAdmin(login string, password string, email string) bool {
	var admin *entitys.Admin
	qb := p.em.pdb.Where(&entitys.Admin{
		Login:    login,
		Pass:     password,
		Email:    email,
		IsActive: true,
	}).First(&admin)

	if util.IsErrQb(qb) || admin.ID == 0 {
		return false
	}

	return true
}

func (p *adminRepo) GetAdmin(login string, password string, email string) *entitys.Admin {
	var admin *entitys.Admin
	qb := p.em.pdb.Where(&entitys.Admin{
		Login:    login,
		Pass:     password,
		Email:    email,
		IsActive: true,
	}).First(&admin)

	if util.IsErrQb(qb) || admin.ID == 0 {
		return nil
	}

	return admin
}

func (p *adminRepo) FindAdmin(id uint) *entitys.Admin {
	var admin *entitys.Admin
	qb := p.em.pdb.Where(&entitys.Admin{
		ID: id,
	}).First(&admin)

	if util.IsErrQb(qb) || admin.ID == 0 {
		return nil
	}

	return admin
}

func (p *adminRepo) FindList(limit int64, offset int64) []*entitys.Admin {
	var data []*entitys.Admin
	qb := p.em.pdb.Where(entitys.Admin{}).Limit(int(limit)).Offset(int(offset)).Find(&data).Order("id desc")

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *adminRepo) FindListCount() int64 {
	var count int64
	var data *[]entitys.Admin
	qb := p.em.pdb.Where(entitys.Admin{}).Find(&data).Count(&count)

	if util.IsErrQb(qb) {
		return 0
	}
	return count
}

func (p *adminRepo) SetByID(
	login string,
	name string,
	email string,
	pass string,
	isActive bool,
	role int64,
	protection int64,
) bool {
	now := time.Now()
	var admin *entitys.Admin
	qb := p.em.pdb.Where(&entitys.Admin{Login: login}).First(&admin)
	admin.Name = name
	admin.Email = email
	admin.IsActive = isActive
	admin.Role = role
	admin.Protection = protection
	admin.UpdateAt = now

	if len(pass) > 0 {
		admin.Pass = pass
	}

	if qb.Error != nil {
		if len(pass) == 0 {
			return false
		}

		admin.Login = login
		admin.CreateAt = now
		qb = p.em.pdb.Create(&admin)
		if qb.Error != nil {
			return false
		}

		return true
	}

	qb = p.em.pdb.Save(&admin)
	if qb.Error != nil {
		return false
	}

	return true
}
