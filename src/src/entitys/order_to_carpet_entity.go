package entitys

import "time"

func (OrderToCarpet) TableName() string {
	return "order_to_carpet"
}

type OrderToCarpet struct {
	ID        uint `gorm:"primaryKey"`
	ClientId  uint
	Client    Client `gorm:"foreignKey:ID"`
	ProductId uint
	Product   Product `gorm:"foreignKey:ID"`
	Status    int64
	Price     int64
	Width     string
	Height    string
	Count     int64
	Uuid      string
	PriceUuid string
	PriceMid  string
	PriceId   int64
	Name      string
	UpdateAt  time.Time
	CreateAt  time.Time
}
