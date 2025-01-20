package persist

import (
	"service_api/src/entitys"
	"service_api/src/util"

	pa "bitbucket.org/klaraeng/package_proto/service_api/admin"

	pb "bitbucket.org/klaraeng/package_proto/service_api/client"
	"github.com/go-kratos/kratos/v2/log"
)

type ImagesContext interface {
	Update(data *entitys.Images) bool
	Create(data *entitys.Images) bool
	FindByID(id int64) *entitys.Images
	FindToImagesFirst(productId int64) *pb.Images
	FindToImagesFirstEntity(productId int64) *entitys.Images
	FindToImagesFirstAdmin(productId int64) *pa.Images
	OneToProductIdImages(id int64) []entitys.Images
}

type imagesRepo struct {
	em  *Persist
	log *log.Helper
}

func NewImagesRepo(p *Persist, logger log.Logger) ImagesContext {
	return &imagesRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *imagesRepo) Update(data *entitys.Images) bool {
	if data.ID == 0 {
		qb := p.em.pdb.Create(&data)
		if qb.Error != nil {
			return false
		}
		return true
	}

	qb := p.em.pdb.Save(&data)
	if qb.Error != nil {
		return false
	}
	return true
}

func (p *imagesRepo) Create(data *entitys.Images) bool {
	qb := p.em.pdb.Create(&data)
	if qb.Error != nil {
		return false
	}
	return true
}

func (p *imagesRepo) FindByID(id int64) *entitys.Images {
	var data *entitys.Images
	qb := p.em.pdb.First(&data, id)
	if util.IsErrQb(qb) {
		return nil
	}
	return data
}

func (p *imagesRepo) FindToImagesFirst(productId int64) *pb.Images {
	var images entitys.Images
	qb := p.em.pdb.Where(&entitys.Images{
		ProductId: productId,
		Type:      false,
	}).First(&images)
	if util.IsErrQb(qb) {
		return nil
	}
	return &pb.Images{
		Name: images.Filename,
		Dir:  images.Dir,
		Type: images.Type,
	}
}

func (p *imagesRepo) FindToImagesFirstEntity(productId int64) *entitys.Images {
	var images *entitys.Images
	qb := p.em.pdb.Where(&entitys.Images{
		ProductId: productId,
		Type:      false,
	}).First(&images)
	if util.IsErrQb(qb) {
		return nil
	}
	return images
}

func (p *imagesRepo) FindToImagesFirstAdmin(productId int64) *pa.Images {
	var images entitys.Images
	qb := p.em.pdb.Where(&entitys.Images{
		ProductId: productId,
		Type:      false,
	}).First(&images)
	if util.IsErrQb(qb) {
		return nil
	}
	return &pa.Images{
		Name: images.Filename,
		Dir:  images.Dir,
		Type: images.Type,
	}
}

func (p *imagesRepo) OneToProductIdImages(id int64) []entitys.Images {
	var images []entitys.Images
	qb := p.em.pdb.Where(&entitys.Images{
		ProductId: id,
	}).Find(&images)
	if util.IsErrQb(qb) || len(images) == 0 {
		return nil
	}
	return images
}
