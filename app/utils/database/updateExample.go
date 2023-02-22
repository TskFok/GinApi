package database

import "github.com/TskFok/GinApi/app/utils/model"

func update() {
	db := GetClient()

	db.Model(&model.Test{}).Where("id=?", 1).Update("title", "1")

	db.Model(&model.Test{}).Where("id=?", 1).Updates(model.Test{
		BaseModel: model.BaseModel{},
		Title:     "a",
		Content:   "c",
	})
}
