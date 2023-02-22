package model

type Test struct {
	BaseModel
	Title   string `gorm:"column:title;type:VARCHAR(50);NOT NULL"`
	Content string `gorm:"column:content;type:TEXT;NOT NULL"`
}
