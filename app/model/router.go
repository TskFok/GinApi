package model

import (
	"github.com/TskFok/GinApi/app/tool"
	"github.com/TskFok/GinApi/app/utils/database"
)

type Router struct {
	BaseModel

	Id          uint32 `gorm:"column:id;type:INT(11) UNSIGNED;AUTO_INCREMENT;NOT NULL" json:"id,omitempty"`
	Router      string `gorm:"column:router;type:VARCHAR(255);NOT NULL" json:"router,omitempty"`
	Description string `gorm:"column:description;type:VARCHAR(255);NOT NULL" json:"description,omitempty"`
	Method      string `gorm:"column:method;type:VARCHAR(50);NOT NULL" json:"method,omitempty"`
	CreatorId   uint32 `gorm:"column:creator_id;type:INT(11) UNSIGNED;NOT NULL" json:"creator_id,omitempty"`
	CreatorName string `gorm:"column:creator_name;type:VARCHAR(50);NOT NULL" json:"creator_name,omitempty"`
}

// Create 创建路由
func (*Router) Create(router *Router) (id uint32, err error) {
	db := database.Db.Create(&router)

	if db.Error != nil {
		return 0, db.Error
	}

	return router.Id, nil
}

// Update 修改路由
func (router *Router) Update(condition interface{}) bool {
	db := database.Db.Model(router).Updates(condition)

	if db.Error != nil {
		return false
	}

	return true
}

// Get 获取路由信息
func (*Router) Get(condition interface{}) (router *Router, exists bool) {
	db := database.Db.Where(condition).First(&router)

	if db.Error != nil {
		return nil, false
	}

	return router, true
}

func (*Router) List(page int, size int) (res map[string]interface{}) {
	db := database.Db.Offset(size * (page - 1)).Limit(size).Order("id desc")

	routers := &[]Router{}

	res = tool.Paginate(db, routers)

	return
}
