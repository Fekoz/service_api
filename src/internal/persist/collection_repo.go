package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"

	"github.com/go-kratos/kratos/v2/log"
)

type CollectionContext interface {
	FindWithId(id int64) []entitys.Collection
}

type collectionRepo struct {
	em  *Persist
	log *log.Helper
}

func NewCollectionRepo(p *Persist, logger log.Logger) CollectionContext {
	return &collectionRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *collectionRepo) FindWithId(id int64) []entitys.Collection {
	var collection []entitys.Collection
	qb := p.em.pdb.Where(&entitys.Collection{
		InProductId: id,
	}).Find(&collection)
	if util.IsErrQb(qb) {
		return nil
	}
	return collection
}
