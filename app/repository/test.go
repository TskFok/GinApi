package repository

import (
	"fmt"
	"github.com/TskFok/GinApi/app/model"
	"github.com/TskFok/GinApi/app/tool"
	"github.com/TskFok/GinApi/app/utils/database"
)

type TestEntity struct {
	Title   string
	Content string
}

func Paginate() {
	//批量查询
	db := database.GetClient()

	ress := make([]TestEntity, 10)

	db = db.Model(&model.Test{}).Where("id > ?", "2").Limit(10)
	rrr, ppp := tool.Paginate(db, &ress)
	fmt.Println(rrr, ppp)
}

func Select() {
	//批量查询
	db := database.GetClient()

	ress := make([]TestEntity, 10)

	db.Model(&model.Test{}).Where("id > ?", "2").Limit(10).Find(&ress)

	for _, v := range ress {
		fmt.Println(v)
	}

	//批量查询
	//res, err := db.Model(&model.Test{}).
	//	Select("title", "content").
	//	Where("Title = ? AND Content = ?", "Hello", "").Rows()
	//
	//if nil != err {
	//	panic(err)
	//}
	//
	//defer res.Close()
	//for res.Next() {
	//	aaa := &model.Test{}
	//	res.Scan(&aaa.Title, &aaa.Content)
	//	fmt.Println(aaa.Title)
	//}
}

func Get() {
	db := database.GetClient()

	test := &model.Test{}
	//查询id=2
	result := db.First(test, 2)

	fmt.Println(result)
}
func Create() uint32 {
	db := database.GetClient()
	test := &model.Test{
		BaseModel: model.BaseModel{},
		Title:     "Hello",
	}

	db.Create(test)

	return test.Id
}

func Update() {
	db := database.GetClient()

	db.Model(&model.Test{}).Where("id=?", 1).Update("title", "1")

	db.Model(&model.Test{}).Where("id=?", 1).Updates(model.Test{
		BaseModel: model.BaseModel{},
		Title:     "a",
		Content:   "c",
	})
}
