package orderr

import (
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/models"
)

// DeleteOrderById 透過 id 刪除訂單
func (r *repo) DeleteOrderById(id string) (apiErr errorcode.Error) {

	db, apiErr := r.DB.DBConn()
	if apiErr != nil {
		return
	}

	// 建立訂單
	if err := db.Where("id = ?", id).Delete(&models.Order{}).Error; err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.DeleteOrderFail, err, id)
		return
	}

	return
}
