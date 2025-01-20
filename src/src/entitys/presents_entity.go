package entitys

import "time"

func (Presents) TableName() string {
	return "presents"
}

type Presents struct {
	ID uint `gorm:"primarykey"`
	Key string
	Value string
	Name string
	UpdateAt time.Time
	CreateAt time.Time
}
