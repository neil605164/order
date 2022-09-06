package orderr

import (
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/models"
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
