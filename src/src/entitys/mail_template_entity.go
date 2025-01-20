package entitys

import "time"

func (MailTemplate) TableName() string {
	return "mail_template"
}

type MailTemplate struct {
	ID uint `gorm:"primaryKey"`
	Title string
	Message string
	UpdateAt time.Time
	CreateAt time.Time
}
