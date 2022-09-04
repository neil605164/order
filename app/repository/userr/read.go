package userr

import (
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/models"
)

// UserList 會員清單
func (r *repo) UserList() (data []models.User, apiErr errorcode.Error) {

	db, apiErr := r.DB.DBConn()
	if apiErr != nil {
		return
	}

	if err := db.Preload("Review").Find(&data).Error; err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.GetUserListError, err)
		return
	}
	return
}
