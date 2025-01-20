package entitys

import "time"

func (Admin) TableName() string {
	return "admin"
}

type Admin struct {
	ID			uint `gorm:"primarykey"`
	Login		string
	Pass		string
	Email		string
	Protection	int64
	Role		int64
	IsActive	bool
	Name		string
	UpdateAt	time.Time
	CreateAt	time.Time
}