package entitys

import "time"

func (Sender) TableName() string {
	return "sender"
}

type Sender struct {
	ID uint `gorm:"primaryKey"`
	ClientId uint
	Client Client `gorm:"foreignKey:ID"`
	MailTemplateId uint
	IsSend bool
	Uuid string
	Type int64
	UpdateAt time.Time
	CreateAt time.Time
}