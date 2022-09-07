package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	Id            uint64          `json:"id" gorm:"primaryKey;unsigned;autoIncrement;comment:流水號"`
	ProductID     uint64          `json:"productId" gorm:"column:product_id;comment:產品ID"`
	OrderNo       string          `json:"orderNo" gorm:"column:order_no;type:varchar(255);comment:訂單編號;NOT NULL;index:order"`
	Behavior      string          `json:"behavior" gorm:"column:behavior;type:varchar(255);comment:交易行為;NOT NULL; index:behavior"`
	Price         decimal.Decimal `json:"price" gorm:"column:price;type:decimal(10,2);comment:價格;NOT NULL"`
	Amount        int             `json:"amount" gorm:"column:amount;type:int(8);comment:數量;NOT NULL"`
	AlreadyAmount int             `json:"alreadyAmount" gorm:"column:already_amount;type:int(8);comment:已成交數量;NOT NULL;default:0"`
	Status        string          `json:"status" gorm:"column:status;type:varchar(30);comment:訂單狀態;NOT NULL;index:status;default:unpaid"`
	PaidAt        *time.Time      `json:"payedAt" gorm:"column:paid_at;type:datetime;comment:付費日期"`
	Product       Product         `json:"product" gorm:"foreignKey:ProductID"`
	CreatedAt     time.Time       `json:"created_at" gorm:"type:datetime comment '建立時間';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time       `json:"updated_at" gorm:"type:datetime comment '更新時間';not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt  `gorm:"index;comment:軟刪除時間" json:"-"`
}
