package entitys

import "time"

func (Category) TableName() string {
	return "category"
}

type Category struct {
	ID uint `gorm:"primarykey"`
	Key string
	Type string
	Value string
	IsActive bool
	Name string
	UpdateAt time.Time
	CreateAt time.Time
}
