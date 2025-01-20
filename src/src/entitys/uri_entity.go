package entitys

import "time"

func (Uri) TableName() string {
	return "uri"
}

type Uri struct {
	ID       uint `gorm:"primaryKey"`
	Uid      string
	IsActive bool
	Uri      string
	Type     int64
	UpdateAt time.Time
	CreateAt time.Time
}
