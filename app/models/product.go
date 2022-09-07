package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id   uint64 `json:"id" gorm:"primaryKey;unsigned;autoIncrement;comment:流水號"`
	Name string `json:"name" gorm:"comment:產品名稱;NOT NULL;index:plan"`
	// Amount    int            `json:"count" gorm:"comment:數量;NOT NULL"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:datetime comment '建立時間';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:datetime comment '更新時間';not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:軟刪除時間" json:"-"`
}
