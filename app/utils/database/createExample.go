package database

import (
	"github.com/TskFok/GinApi/app/utils/model"
)

// Create /创建一条记录
func create() uint32 {
	db := GetClient()
	test := &model.Test{
		BaseModel: model.BaseModel{},
		Title:     "Hello",
	}

	db.Create(test)

	return test.Id
}
