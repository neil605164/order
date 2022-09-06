package structer

import (
	"time"

	"github.com/shopspring/decimal"
)

// OrderListResp 回傳API格式
type OrderListResp struct {
	Result []OrderList `json:"result"`
	Status RespStatus  `json:"status"`
}

// OrderDetailResp 回傳API格式
type OrderDetailResp struct {
	Result OrderList  `json:"result"`
	Status RespStatus `json:"status"`
}

type OrderList struct {
	Id        uint64          `json:"id" example:"1"`
	CreatedAt time.Time       `json:"created_at" example:"2022-07-21T12:19:39-04:00"`
	OrderNo   string          `json:"order_no" example:"qa12sw43vr345"`
	ProductID uint64          `json:"productId" example:"1"`
	Behavior  string          `json:"behavior" example:"sell"`
	Price     decimal.Decimal `json:"price" example:"100"`
	Amount    int             `json:"amount" example:"10"`
	Status    string          `json:"status" example:"unpaid"`
	PaidAt    *time.Time      `json:"payedAt"  example:"2022-07-21T12:19:39-04:00"`
	Product   Product         `json:"product"`
}

type OrderDetail struct {
	Id        uint64    `json:"id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2022-07-21T12:19:39-04:00"`
}

type Product struct {
	Id     uint64 `json:"id" example:"1"`
	Name   string `json:"name" example:"產品名稱一"`
	Amount int    `json:"count" example:"100"`
}
