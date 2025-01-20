package entitys

import "time"

func (Viewcases) TableName() string {
	return "viewcases"
}

type Viewcases struct {
	ID       uint `gorm:"primarykey"`
	Uuid     string
	IsActive bool
	List     string `sql:"type:jsonb"`
	Name     string
	UpdateAt time.Time
	CreateAt time.Time
}
