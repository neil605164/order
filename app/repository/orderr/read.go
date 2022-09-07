package orderr

import (
	"errors"
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/models"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

// GetOrderList 取 order list 資料
func (r *repo) GetOrderList() (orders []models.Order, apiErr errorcode.Error) {

	db, apiErr := r.DB.DBConn()
	if apiErr != nil {
		return
	}

	if err := db.Preload("Product").Find(&orders).Error; err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.GetOrderListError, err)
		return
	}

	return
}

// GetOrderById 透過 id 取 order 資料
func (r *repo) GetOrderById(id string) (order models.Order, apiErr errorcode.Error) {

	db, apiErr := r.DB.DBConn()
	if apiErr != nil {
		return
	}

	if err := db.Preload("Product").Where("id = ?", id).Find(&order).Error; err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.GetOrderByIdError, err)
		return
	}

	return
}

// GetOrderByBehaviorAndPrice 透過行為與價格進行撮合
func (r *repo) GetOrderByBehaviorAndPrice(behavior string, price decimal.Decimal, productId uint64) (orders []models.Order, apiErr errorcode.Error) {
	db, apiErr := r.DB.DBConn()
	if apiErr != nil {
		return
	}

	// 取尚未交易成功的訂單
	err := db.
		Where("behavior = ? AND price = ? AND status =? AND product_id = ?", behavior, price, global.OrderUnpaid, productId).
		Find(&orders).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.GetOrderListError, err)
		return
	}
	return
}
