package models

import (
	"time"

	"gorm.io/gorm"
)

type Id struct {
	Id uint64 `json:"id" gorm:"primaryKey;unsigned;autoIncrement;comment:流水號"`
}
type BaseTime struct {
	CreatedAt time.Time      `json:"created_at" gorm:"type:datetime comment '建立時間';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:datetime comment '更新時間';not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:軟刪除時間" json:"-"`
}
