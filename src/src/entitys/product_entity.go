package entitys

import (
	"time"
)

func (Product) TableName() string {
	return "product"
}

type Product struct {
	ID int64 `gorm:"primarykey"`
	IsActive bool
	Article string
	Uuid string
	FullPrice string
	Factor int64
	Name string
	OriginalUrl string
	// Locker
	IsProductLock bool
	IsPriceLock bool
	IsSpecificationLock bool
	IsImagesLock bool
	IsAttributeLock bool
	IsGlobalUpdateLock bool
	// Locker
	UpdateAt time.Time
	CreateAt time.Time
}
