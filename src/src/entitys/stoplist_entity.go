package entitys

import "time"

func (StopList) TableName() string {
	return "stop_list"
}

type StopList struct {
	ID uint `gorm:"primarykey"`
	Ip string
	Type int64
	Session string
	UnlockAt time.Time
	Try int64
	Name string
	UpdateAt time.Time
	CreateAt time.Time
}