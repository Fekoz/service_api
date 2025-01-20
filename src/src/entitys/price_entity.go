package entitys

func (Price) TableName() string {
	return "price"
}

type Price struct {
	ID uint `gorm:"primarykey"`
	Width string
	Height string
	Count int64
	Price float64
	Meter int64
	OldPrice float64
	Uuid string
	Mid string
	Storage int64
	ProductId int64
}
