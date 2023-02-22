package tool

import (
	"gorm.io/gorm"
)

func Paginator(tx *gorm.DB, desc any) (any, int64) {
	tx.Find(desc)
	var count int64
	tx.Count(&count)

	return desc, count
}
