package orderr

import (
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/models"
)

// UpdateOrderById 透過 id 更新訂單
func (r *repo) UpdateOrderById(id string, order map[string]interface{}) (apiErr errorcode.Error) {

	db, apiErr := r.DB.DBConn()
	if apiErr != nil {
		return
	}

	if err := db.Model(&models.Order{}).Updates(order).Error; err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.UpdateOrderFail, err, id)
		return
	}

	return
}
