package entitys

func (Collection) TableName() string {
	return "collection"
}

type Collection struct {
	ID           uint `gorm:"primarykey"`
	InProductId  int64
	OutProductId int64
}
