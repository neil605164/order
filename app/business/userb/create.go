package userb

import (
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/global/structer"
)

// CreateUser 創立 user
func (b *business) CreateUser(raw *structer.CreateReq) (apiErr errorcode.Error) {

	// 密碼加密
	raw.Pwd, apiErr = helper.HashPassword(raw.Password)
	if apiErr != nil {
		return
	}

	// 亂數產生 user memberNo
	raw.MemberNo = helper.RandStringBytesMaskImpr(10)

	// 建立用戶
	if apiErr = b.user.CreateUser(raw); apiErr != nil {
		return
	}
	return
}
