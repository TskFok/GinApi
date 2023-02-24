package model

type Test struct {
	BaseModel
	Title   string `json:"title" gorm:"column:title;type:VARCHAR(50);NOT NULL"`
	Content string `json:"content" gorm:"column:content;type:TEXT;NOT NULL"`
}

func GetTest(filter interface{}) (test []Test) {
	db.Model(&Test{}).Where(filter).Find(&test)

	return
}
