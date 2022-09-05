package orderr

import (
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
)

// CreateOrder 建立訂單
func (r *repo) CreateOrder(order map[string]interface{}) (apiErr errorcode.Error) {
	db, apiErr := r.DB.DBConn()
	if apiErr != nil {
		return
	}

	// 建立訂單
	if err := db.Create(&order).Error; err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.CreateUserFail, err, order)
		return
	}

	return nil
}
