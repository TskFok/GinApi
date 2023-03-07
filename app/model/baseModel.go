package model

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:DATETIME;NOT NULL"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:DATETIME;NOT NULL"`
}
