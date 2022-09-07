package orderb

import (
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/global/structer"
	"order/app/models"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

// GetOrderById 透過 id 取 order 資料
func (b *business) GetOrderList() (resp []structer.OrderList, apiErr errorcode.Error) {
	// 取 db 資料
	orders, apiErr := b.order.GetOrderList()
	if apiErr != nil {
		return
	}

	// 組合資料
	resp = []structer.OrderList{}
	for k := range orders {
		orderTmp := structer.OrderList{}
		orderTmp.Id = orders[k].Id
		orderTmp.OrderNo = orders[k].OrderNo
		orderTmp.ProductID = orders[k].ProductID
		orderTmp.Behavior = orders[k].Behavior
		orderTmp.Price = orders[k].Price
		orderTmp.Amount = orders[k].Amount
		orderTmp.Status = orders[k].Status
		orderTmp.PaidAt = orders[k].PaidAt
		orderTmp.Product = structer.Product{}

		orderTmp.Product.Id = orders[k].Product.Id
		orderTmp.Product.Name = orders[k].Product.Name

		resp = append(resp, orderTmp)
	}

	return
}

// GetOrderById 透過 id 取 order 資料
func (b *business) GetOrderById(id string) (resp *structer.OrderList, apiErr errorcode.Error) {
	// 取 db 資料
	order, apiErr := b.order.GetOrderById(id)
	if apiErr != nil {
		return
	}

	if order.Id != 0 {
		// 組合資料
		resp = &structer.OrderList{}
		resp.Id = order.Id
		resp.OrderNo = order.OrderNo
		resp.ProductID = order.ProductID
		resp.Behavior = order.Behavior
		resp.Price = order.Price
		resp.Amount = order.Amount
		resp.Status = order.Status
		resp.PaidAt = order.PaidAt
		resp.Product = structer.Product{}
		resp.Product.Id = order.Product.Id
		resp.Product.Name = order.Product.Name
	}

	return
}

// MatchOrder 搓合訂單
func (b *business) MatchOrder(byteData []byte) (apiErr errorcode.Error) {

	order := models.Order{}
	if err := jsoniter.Unmarshal(byteData, &order); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.JSONUnMarshalError, nil, string(byteData))
		return
	}

	switch order.Behavior {
	case global.OrderBuyBehavior:
		// 至 DB 取賣單進行撮合
		dbOrders, apiErr := b.order.GetOrderByBehaviorAndPrice(global.OrderSellBehavior, order.Price, order.ProductID)
		if apiErr != nil {
			return apiErr
		}

		if apiErr := b.matchOrder(dbOrders, int(order.Id)); apiErr != nil {
			return apiErr
		}
	case global.OrderSellBehavior:
		// 至 DB 取買單進行撮合
		dbOrders, apiErr := b.order.GetOrderByBehaviorAndPrice(global.OrderBuyBehavior, order.Price, order.ProductID)
		if apiErr != nil {
			return apiErr
		}

		if apiErr := b.matchOrder(dbOrders, int(order.Id)); apiErr != nil {
			return apiErr
		}
	}

	return
}

func (b *business) matchOrder(dbOrders []models.Order, orderId int) (apiErr errorcode.Error) {

	id := strconv.Itoa(orderId)

	for k := range dbOrders {
		// 取新訂單即時狀態
		order, err := b.order.GetOrderById(id)
		if err != nil {
			return
		}

		alreadyAmount := order.Amount - order.AlreadyAmount

		// 數量已達標將不在進行撮合
		if alreadyAmount <= 0 {
			return
		}

		now := time.Now()

		switch alreadyAmount > 0 {
		// 如果 db 訂單數量大於需求數量
		case (dbOrders[k].Amount - dbOrders[k].AlreadyAmount) > alreadyAmount:
			// 更新 db 訂單
			id := strconv.Itoa(int(dbOrders[k].Id))
			orderMap := make(map[string]interface{})
			orderMap["paid_at"] = now
			orderMap["already_amount"] = dbOrders[k].AlreadyAmount + alreadyAmount

			if orderMap["already_amount"] == dbOrders[k].Amount {
				orderMap["status"] = global.OrderSuccess
			}

			b.order.UpdateOrderById(id, orderMap)
		// 如果 db 訂單數量小於需求數量
		case (dbOrders[k].Amount - dbOrders[k].AlreadyAmount) < alreadyAmount:
			id := strconv.Itoa(int(dbOrders[k].Id))
			orderMap := make(map[string]interface{})
			orderMap["paid_at"] = now
			orderMap["already_amount"] = dbOrders[k].Amount
			orderMap["status"] = global.OrderSuccess

			b.order.UpdateOrderById(id, orderMap)

			// 扣除已匹配的數量
			alreadyAmount = (dbOrders[k].Amount - dbOrders[k].AlreadyAmount)
		// 如果 db 訂單數量等於需求數量
		case (dbOrders[k].Amount - dbOrders[k].AlreadyAmount) == alreadyAmount:
			// 更新 db 訂單
			id := strconv.Itoa(int(dbOrders[k].Id))
			orderMap := make(map[string]interface{})
			orderMap["status"] = global.OrderSuccess
			orderMap["paid_at"] = now

			// 檢查是否為最後一單成交量
			orderMap["already_amount"] = dbOrders[k].AlreadyAmount + alreadyAmount
			if dbOrders[k].AlreadyAmount+alreadyAmount != dbOrders[k].Amount {
				orderMap["already_amount"] = alreadyAmount
				alreadyAmount = dbOrders[k].Amount
			}

			b.order.UpdateOrderById(id, orderMap)

			// 扣除已匹配的數量

		}

		// 更新新訂單
		id := strconv.Itoa(int(order.Id))
		orderMap := make(map[string]interface{})
		// 檢查是否為最後一單成交量
		orderMap["already_amount"] = order.AlreadyAmount + alreadyAmount
		if order.AlreadyAmount+alreadyAmount != order.Amount {
			orderMap["already_amount"] = alreadyAmount
		}
		orderMap["paid_at"] = now

		if orderMap["already_amount"] == order.Amount {
			orderMap["status"] = global.OrderSuccess
		}

		b.order.UpdateOrderById(id, orderMap)

		// 如果該筆訂單已經達到撮合數量，直接離開 for loop
		if order.Amount == alreadyAmount {
			break
		}
	}
	return
}
