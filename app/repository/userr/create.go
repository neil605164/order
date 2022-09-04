package userr

import (
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/global/structer"
	"order/app/models"
)

// CreateUser 建立用戶
func (r *repo) CreateUser(raw *structer.CreateReq) (apiErr errorcode.Error) {
	db, apiErr := r.DB.DBConn()
	if apiErr != nil {
		return
	}

	// 組資料
	user := models.User{
		MemberNo: raw.MemberNo,
		Username: raw.Username,
		Email:    raw.Email,
		Password: raw.Pwd,
		Birthday: raw.Birthday,
		Review:   &models.UserReview{},
	}

	user.Review = &models.UserReview{
		Status: global.NotVerify,
	}

	tx := db.Begin()

	// 建立會員
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.CreateUserFail, err, user)
		return
	}

	// 更新會員資料
	mapData := map[string]interface{}{}
	mapData["review_id"] = user.Review.Id

	if err := tx.Model(models.User{}).Where("id = ?", user.Id).Updates(mapData).Error; err != nil {
		tx.Rollback()
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.UpdateUserFail, err, user)
		return
	}

	tx.Commit()

	return nil
}
