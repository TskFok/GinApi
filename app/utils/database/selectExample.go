package database

import (
	"fmt"
	"github.com/TskFok/GinApi/app/utils/model"
)

func sselect() {
	db := GetClient()

	test := &model.Test{}
	//查询id=2
	result := db.First(test, 2)

	fmt.Println(test.Title, result.RowsAffected)

	//批量查询
	res, err := db.Model(&model.Test{}).
		Select("title", "content").
		Where("Title = ? AND Content = ?", "Hello", "").Rows()

	if nil != err {
		panic(err)
	}

	defer res.Close()
	for res.Next() {
		aaa := &model.Test{}
		res.Scan(&aaa.Title, &aaa.Content)
		fmt.Println(aaa.Title)
	}

}
