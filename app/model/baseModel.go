package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)

	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t *LocalTime) Value() (driver.Value, error) {
	return time.Time(*t), nil
}

type BaseModel struct {
	CreatedAt *LocalTime `json:"created_at" gorm:"column:created_at;type:DATETIME;NOT NULL"`
	UpdatedAt *LocalTime `json:"updated_at" gorm:"column:updated_at;type:DATETIME;NOT NULL"`
}
