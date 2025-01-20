package persist

import (
	"database/sql"
	"service_api/src/dtos"

	"github.com/go-kratos/kratos/v2/log"
)

type ReportsContext interface {
	FindByMarketPlace(minCount int64, limit int64, offset int64) []dtos.MarketTypesTransfer
	FindByMarketPlaceCount(minCount int64, limit int64, offset int64) int64
}

type reportsRepo struct {
	em  *Persist
	log *log.Helper
}

func NewReportsRepo(p *Persist, logger log.Logger) ReportsContext {
	return &reportsRepo{
		em:  p,
		log: log.NewHelper(logger),
	}
}

func (p *reportsRepo) FindByMarketPlaceCount(minCount int64, limit int64, offset int64) int64 {
	var totalCount int64 = 0
	p.em.pdb.Table("price as s").
		Select("distinct p.article as article, s.width as width, s.height as height, 0 + COALESCE(s0.count, 0) + COALESCE(s1.count, 0) + COALESCE(s2.count, 0) > 2 as total_count, s.mid as mid, u.type as type_market, u.uri as url, CASE WHEN u.type = 1 THEN concat('name.', s.mid) WHEN u.type is null THEN null ELSE s.mid END as marketplace").
		Joins("JOIN product p ON s.product_id = p.id AND p.is_active = true").
		Joins("left JOIN uri u ON u.uid = s.mid").
		Joins("left JOIN price s0 ON s0.mid = s.mid and s0.storage = 4").
		Joins("left JOIN price s1 ON s1.mid = s.mid and s1.storage = 0").
		Joins("left JOIN price s2 ON s2.mid = s.mid and s2.storage = 7").
		Where("p.is_active = true AND 0 + COALESCE(s0.count, 0) + COALESCE(s1.count, 0) + COALESCE(s2.count, 0) > @count AND p.is_active = true", sql.Named("count", minCount)).
		Count(&totalCount)
	return totalCount
}

func (p *reportsRepo) FindByMarketPlace(minCount int64, limit int64, offset int64) []dtos.MarketTypesTransfer {
	var results []dtos.MarketTypesTransfer
	p.em.pdb.Table("price as s").
		Select("distinct p.article as article, s.width as width, s.height as height, 0 + COALESCE(s0.count, 0) + COALESCE(s1.count, 0) + COALESCE(s2.count, 0) > 2 as total_count, s.mid as mid, u.type as type_market, u.uri as url, CASE WHEN u.type = 1 THEN concat('name.', s.mid) WHEN u.type is null THEN null ELSE s.mid END as marketplace").
		Joins("JOIN product p ON s.product_id = p.id AND p.is_active = true").
		Joins("left JOIN uri u ON u.uid = s.mid").
		Joins("left JOIN price s0 ON s0.mid = s.mid and s0.storage = 4").
		Joins("left JOIN price s1 ON s1.mid = s.mid and s1.storage = 0").
		Joins("left JOIN price s2 ON s2.mid = s.mid and s2.storage = 7").
		Where("p.is_active = true AND 0 + COALESCE(s0.count, 0) + COALESCE(s1.count, 0) + COALESCE(s2.count, 0) > @count AND p.is_active = true", sql.Named("count", minCount)).
		Limit(int(limit)).Offset(int(offset)).
		Scan(&results)
	return results
}
