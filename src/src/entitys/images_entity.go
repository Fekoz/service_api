package entitys

func (Images) TableName() string {
	return "images"
}

type Images struct {
	ID uint `gorm:"primarykey"`
	Dir string
	Filename string
	OriginalUrl string
	Type bool
	ProductId int64
}

