package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"
	"time"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/log"
)

type ClientContext interface {
	FindListCount() int64
	FindList(limit int64, offset int64) []*entitys.Client
	CreateOrUpdateClient(request *pb.Personal) (uint, error)
	CreateOrUpdateClientReturnClient(request *pb.Personal) (*entitys.Client, error)
}

type clientRepo struct {
	em  *Persist
	log *log.Helper
}

func NewClientRepo(p *Persist, logger log.Logger) ClientContext {
	return &clientRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *clientRepo) CreateOrUpdateClient(person *pb.Personal) (uint, error) {
	now := time.Now()
	var client *entitys.Client
	qb := p.em.pdb.Where(&entitys.Client{
		Phone: person.Phone,
		Email: person.Email,
		Name:  person.Name,
	}).First(&client)

	client.UpdateAt = now
	if qb.Error != nil {
		client.Phone = person.Phone
		client.Email = person.Email
		client.Name = person.Name
		client.CreateAt = now
		qb = p.em.pdb.Create(&client)
		if qb.Error != nil {
			return 0, qb.Error
		}
		return client.ID, nil
	}

	qb = p.em.pdb.Save(&client)
	if qb.Error != nil {
		return 0, qb.Error
	}
	return client.ID, nil
}

func (p *clientRepo) CreateOrUpdateClientReturnClient(person *pb.Personal) (*entitys.Client, error) {
	now := time.Now()
	var client *entitys.Client
	qb := p.em.pdb.Where(&entitys.Client{
		Phone: person.Phone,
		Email: person.Email,
		Name:  person.Name,
	}).First(&client)

	client.UpdateAt = now
	if qb.Error != nil {
		client.Phone = person.Phone
		client.Email = person.Email
		client.Name = person.Name
		client.CreateAt = now
		qb = p.em.pdb.Create(&client)
		if qb.Error != nil {
			return nil, qb.Error
		}
		return client, nil
	}

	qb = p.em.pdb.Save(&client)
	if qb.Error != nil {
		return nil, qb.Error
	}
	return client, nil
}

func (p *clientRepo) FindList(limit int64, offset int64) []*entitys.Client {
	var data []*entitys.Client
	qb := p.em.pdb.Where(entitys.Client{}).Limit(int(limit)).Offset(int(offset)).Find(&data).Order("id desc")

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *clientRepo) FindListCount() int64 {
	var count int64
	var data *[]entitys.Client
	qb := p.em.pdb.Where(entitys.Client{}).Find(&data).Count(&count)

	if util.IsErrQb(qb) {
		return 0
	}
	return count
}
