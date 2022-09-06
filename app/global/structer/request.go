package structer

import (
	"github.com/shopspring/decimal"
)

type CreateOrderReq struct {
	ProductID uint64          `json:"productId" validate:"required" example:"1"`
	Behavior  string          `json:"behavior" validate:"required,oneof=buy sell" example:"buy"`
	Price     decimal.Decimal `json:"price" validate:"required" example:"5"`
	Amount    int             `json:"amount" validate:"required" example:"10"`
}

// type UpdateOrderReq struct {
// 	ProductID uint64          `json:"productId" validate:"required" example:"1"`
// 	Behavior  string          `json:"behavior" validate:"required,oneof=buy sell" example:"buy"`
// 	Price     decimal.Decimal `json:"price" validate:"required" example:"5"`
// 	Amount    int             `json:"amount" validate:"required" example:"10"`
// }
