package model

import (
	"time"
)

type BaseModel struct {
	Id        uint32    `json:"id" gorm:"column:id;type:INT(11) UNSIGNED;AUTO_INCREMENT;NOT NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:DATETIME;NOT NULL"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:DATETIME;NOT NULL"`
}
