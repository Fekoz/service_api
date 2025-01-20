package entitys

import "time"

func (PriceExportDynamic) TableName() string {
	return "price_export_dynamic"
}

type PriceExportDynamic struct {
	ID            uint `gorm:"primarykey"`
	MinWidth      int64
	MaxWidth      int64
	MinHeight     int64
	MaxHeight     int64
	PackageWidth  int64
	PackageHeight int64
	PackageDepth  int64
	UpdateAt      time.Time
	CreateAt      time.Time
}
