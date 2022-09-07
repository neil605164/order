package orderb

import (
	"fmt"
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/global/structer"
	"order/app/models"
	"time"

	jsoniter "github.com/json-iterator/go"
)

// CreateUser 創立 user
func (b *business) CreateOrder(raw structer.CreateOrderReq) (apiErr errorcode.Error) {
	// 檢查金額是否合法
	if raw.Price.IsNegative() || raw.Price.IsZero() {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.IllegalAmount, nil, raw)
		return
	}

	// 取產品資料
	isExist, apiErr := b.product.CheckProudctExistsById(raw.ProductID)
	if apiErr != nil {
		return
	}

	if !isExist {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.ProductNotExist, nil, raw)
		return
	}

	// 產生訂單編號
	orderNo := fmt.Sprintf("%v", time.Now().Unix()) + helper.RandString(6)

	// 組資料
	order := models.Order{
		ProductID: raw.ProductID,
		OrderNo:   orderNo,
		Behavior:  raw.Behavior,
		Price:     raw.Price,
		Amount:    raw.Amount,
		Status:    global.OrderUnpaid,
	}

	// 寫入 order db
	order, apiErr = b.order.CreateOrder(order)
	if apiErr != nil {
		return
	}

	// 處理丟入 queue 資料
	byteData, err := jsoniter.Marshal(order)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.JSONMarshalError, nil, raw)
		return
	}
	queueData := structer.RedisLPushFormat{
		Type: global.OrderQueue,
		Data: order,
	}

	byteData, err = jsoniter.Marshal(queueData)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.JSONMarshalError, nil, raw)
		return
	}

	// 丟入 queue 進行撮合
	if err := b.cache.LPush(global.RedisQueueChannel, byteData); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.OrderPushIntoQueueFail, nil, string(byteData))
		return
	}

	return
}
