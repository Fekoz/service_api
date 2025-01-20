package entitys

import "time"

func (Client) TableName() string {
	return "client"
}

type Client struct {
	ID uint `gorm:"primarykey"`
	Phone string
	Email string
	Name string
	UpdateAt time.Time
	CreateAt time.Time
}
