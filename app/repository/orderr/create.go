package orderr

import (
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/models"
)

// CreateOrder 建立訂單
func (r *repo) CreateOrder(order models.Order) (models.Order, errorcode.Error) {
	db, apiErr := r.DB.DBConn()
	if apiErr != nil {
		return models.Order{}, apiErr
	}

	// 建立訂單
	if err := db.Create(&order).Error; err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.CreateOrderFail, err, order)
		return models.Order{}, apiErr
	}

	return order, nil
}
