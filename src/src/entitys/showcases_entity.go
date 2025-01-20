package entitys

import "time"

func (Showcases) TableName() string {
	return "showcases"
}

type Showcases struct {
	ID uint `gorm:"primarykey"`
	Uuid string
	IsActive bool
	PriceList string `sql:"type:jsonb"`
	Name string
	UpdateAt time.Time
	CreateAt time.Time
}

