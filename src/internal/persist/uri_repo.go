package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"

	"github.com/go-kratos/kratos/v2/log"
)

type UriContext interface {
	FindByMIDs(uid []string) []entitys.Uri
}

type uriRepo struct {
	em  *Persist
	log *log.Helper
}

func NewUriRepo(p *Persist, logger log.Logger) UriContext {
	return &uriRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *uriRepo) FindByMIDs(uid []string) []entitys.Uri {
	var data []entitys.Uri
	qb := p.em.pdb.Where("uid IN ?", uid).Where("is_active = ?", true).Find(&data)

	if util.IsErrQb(qb) {
		return nil
	}
	return data
}
