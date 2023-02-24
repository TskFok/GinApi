package model

import "fmt"

type Test struct {
	BaseModel
	Title   string `json:"title" gorm:"column:title;type:VARCHAR(50);NOT NULL"`
	Content string `json:"content" gorm:"column:content;type:TEXT;NOT NULL"`
}

func (ttt *Test) GetTest(filter interface{}) (test []Test) {
	db.Model(&Test{}).Where(filter).Find(&test)

	return
}

func (ttt *Test) GetOne(id string) (test Test) {
	db.Where("id = ?", id).First(&test)

	return
}

func (ttt *Test) Update(filter interface{}) bool {
	fmt.Println(ttt.Id, filter)
	db = db.Model(ttt).Updates(filter)

	if nil != db.Error {
		return false
	}
	return true
}
