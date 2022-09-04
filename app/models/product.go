package models

type Product struct {
	Id
	Name   string `json:"name" gorm:"comment:產品名稱;NOT NULL;index:plan"`
	Amount int    `json:"count" gorm:"comment:數量;NOT NULL"`
	BaseTime
}
