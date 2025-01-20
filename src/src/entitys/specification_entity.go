package entitys

func (Specification) TableName() string {
	return "specification"
}

type Specification struct {
	ID uint `gorm:"primarykey"`
	Name string
	Value string
	ProductId int64
}
