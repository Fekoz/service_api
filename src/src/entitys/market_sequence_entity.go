package entitys

import "time"

func (MarketSequence) TableName() string {
	return "market_sequence"
}

type MarketSequence struct {
	ID         uint `gorm:"primaryKey"`
	Mid        string
	IsActive   bool
	Name       string
	IsDisabled bool
	IsCounter  bool
	CounterPkg int64
	UpdateAt   time.Time
	CreateAt   time.Time
}
