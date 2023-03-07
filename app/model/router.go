package model

import "github.com/TskFok/GinApi/app/utils/database"

type Router struct {
	BaseModel

	Id          uint32 `gorm:"column:id;type:INT(11) UNSIGNED;AUTO_INCREMENT;NOT NULL"`
	Router      string `gorm:"column:router;type:VARCHAR(255);NOT NULL"`
	Description string `gorm:"column:description;type:VARCHAR(255);NOT NULL"`
	Type        string `gorm:"column:type;type:VARCHAR(50);NOT NULL"`
	CreatorId   uint32 `gorm:"column:creator_id;type:INT(11) UNSIGNED;NOT NULL"`
	CreatorName string `gorm:"column:creator_name;type:VARCHAR(50);NOT NULL"`
}

func (router *Router) Create(newRouter *Router) (id uint32, err error) {
	db := database.Db.Create(&newRouter)

	if db.Error != nil {
		return 0, db.Error
	}

	return newRouter.Id, nil
}
