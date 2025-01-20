package entitys

import "time"

func (Order) TableName() string {
	return "order"
}

type Order struct {
	ID uint `gorm:"primaryKey"`
	ClientId uint
	Client Client `gorm:"foreignKey:ID"`
	PresentsId uint
	Presents Presents `gorm:"foreignKey:ID"`
	Status int64
	Type int64
	Name string
	Uuid string
	UpdateAt time.Time
	CreateAt time.Time
}
