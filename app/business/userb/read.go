package userb

import (
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/global/structer"

	jsoniter "github.com/json-iterator/go"
)

// UserList 使用者清單
func (b *business) UserList() ([]structer.UserList, errorcode.Error) {
	resp := []structer.UserList{}
	// 取 DB 資料
	users, apiErr := b.user.UserList()
	if apiErr != nil {
		return resp, apiErr
	}

	// 組合資料
	for k := range users {
		tmp := structer.UserList{}

		byteData, err := jsoniter.Marshal(users[k])
		if err != nil {
			apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.JSONMarshalError, err, users)
			return resp, apiErr
		}

		if err := jsoniter.Unmarshal(byteData, &tmp); err != nil {
			apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.JSONUnMarshalError, err, users)
			return resp, apiErr
		}

		resp = append(resp, tmp)
	}

	return resp, apiErr
}
