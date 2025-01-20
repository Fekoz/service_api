package entitys

func (Attribute) TableName() string {
	return "attribute"
}

type Attribute struct {
	ID uint `gorm:"primarykey"`
	Name string
	Value string
	ProductId int64
}
