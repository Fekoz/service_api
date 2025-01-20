package entitys

func (OrderField) TableName() string {
	return "order_field"
}

type OrderField struct {
	ID uint `gorm:"primaryKey;"`
	OrderID uint `gorm:"foreignKey:ID"`
	Key string
	Value string
}

