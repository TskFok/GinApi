package tool

import (
	"gorm.io/gorm"
)

func Paginate(tx *gorm.DB, desc any) map[string]interface{} {
	tx.Find(desc)
	var count int64
	tx.Count(&count)

	res := make(map[string]interface{})
	res["data_list"] = desc
	res["count"] = count

	return res
}
