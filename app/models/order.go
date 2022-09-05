package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	Id
	ProductID uint64          `json:"productId" gorm:"column:product_id;comment:產品ID"`
	OrderNo   string          `json:"orderNo" gorm:"column:order_no;type:varchar(255);comment:訂單編號;NOT NULL;index:order"`
	Behavior  string          `json:"behavior" gorm:"column:behavior;type:varchar(255);comment:交易行為;NOT NULL; index:behavior"`
	Price     decimal.Decimal `json:"price" gorm:"column:price;type:decimal(10,2);comment:價格"`
	Amount    int             `json:"amount" gorm:"column:amount;type:decimal(10,2);comment:數量"`
	Status    string          `json:"status" gorm:"column:status;type:varchar(30);comment:訂單狀態;NOT NULL;index:status;default:unpaid"`
	PaidAt    *time.Time      `json:"payedAt" gorm:"column:paid_at;type:datetime;comment:付費日期"`
	Product   Product         `json:"product" gorm:"foreignKey:ProductID"`
	BaseTime
}
