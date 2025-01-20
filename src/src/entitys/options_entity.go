package entitys

import "time"

func (Options) TableName() string {
	return "options"
}

type Options struct {
	ID uint `gorm:"primarykey"`
	Name string
	Value string
	Info string
	UpdateAt time.Time
	CreateAt time.Time
}