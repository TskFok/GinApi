package model

import "time"

type BaseModel struct {
	Id        uint32    `gorm:"column:id;type:INT(11) UNSIGNED;AUTO_INCREMENT;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;type:DATETIME;NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:DATETIME;NOT NULL"`
}
