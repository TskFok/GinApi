package model

import (
	"github.com/TskFok/GinApi/app/utils/database"
	"github.com/TskFok/GinApi/app/utils/logger"
	"time"
)

type User struct {
	BaseModel

	Id            uint32    `gorm:"column:id;type:INT(11) UNSIGNED;AUTO_INCREMENT;NOT NULL" json:"id,omitempty"`
	Nick          string    `gorm:"column:nick;type:VARCHAR(50);NOT NULL" json:"nick,omitempty"`
	UserName      string    `gorm:"column:user_name;type:VARCHAR(50);NOT NULL" json:"user_name,omitempty"`
	Password      string    `gorm:"column:password;type:VARCHAR(255);NOT NULL" json:"password,omitempty"`
	Salt          string    `gorm:"column:salt;type:VARCHAR(255);NOT NULL" json:"salt,omitempty"`
	LastLoginTime time.Time `gorm:"column:last_login_time;type:DATETIME;NOT NULL" json:"last_login_time"`
	LoginIp       string    `gorm:"column:login_ip;type:VARCHAR(255);NOT NULL" json:"login_ip,omitempty"`
}

// Get 获取用户信息
func (*User) Get(condition interface{}) (u *User, exists bool) {
	db := database.Db.Where(condition).First(&u)

	if nil != db.Error {
		logger.Error(db.Error.Error())

		return u, false
	}

	return u, true
}

// Create 创建用户
func (*User) Create(param *User) (uint32, bool) {
	db := database.Db.Create(param)

	if nil != db.Error {
		logger.Error(db.Error.Error())

		return 0, false
	}

	return param.Id, true
}

// Update 更新用户
func (user *User) Update(condition interface{}) bool {
	db := database.Db.Model(user).Updates(condition)

	if nil != db.Error {
		logger.Error(db.Error.Error())

		return false
	}

	return true
}
