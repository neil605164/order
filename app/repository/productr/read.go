package productr

import (
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/models"
)

// GetProductById 透過產品 id 取產品資料
func (r *repo) GetProductById(id uint64) (product models.Product, apiErr errorcode.Error) {
	db, apiErr := r.DB.DBConn()
	if apiErr != nil {
		return
	}

	if err := db.Where("id = ?", id).Find(&product).Error; err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.GetProductByIdError, err)
		return
	}
	return
}

// CheckProudctExistsById 檢查產品是否存在
func (r *repo) CheckProudctExistsById(id uint64) (isExist bool, apiErr errorcode.Error) {
	db, apiErr := r.DB.DBConn()
	if apiErr != nil {
		return
	}

	if err := db.Model(&models.Product{}).Select("count(*) > 0").Where("id = ?", id).Find(&isExist).Error; err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.GetProductByIdError, err)
		return
	}

	return
}
