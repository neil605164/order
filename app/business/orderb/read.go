package orderb

import (
	"order/app/global/errorcode"
	"order/app/global/structer"
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
		orderTmp.Id = orders[k].Id.Id
		orderTmp.OrderNo = orders[k].OrderNo
		orderTmp.ProductID = orders[k].ProductID
		orderTmp.Behavior = orders[k].Behavior
		orderTmp.Price = orders[k].Price
		orderTmp.Amount = orders[k].Amount
		orderTmp.Status = orders[k].Status
		orderTmp.PaidAt = orders[k].PaidAt
		orderTmp.Product = structer.Product{}

		orderTmp.Product.Id = orders[k].Product.Id.Id
		orderTmp.Product.Name = orders[k].Product.Name
		orderTmp.Product.Amount = orders[k].Product.Amount

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

	if order.Id.Id != 0 {
		// 組合資料
		resp = &structer.OrderList{}
		resp.Id = order.Id.Id
		resp.OrderNo = order.OrderNo
		resp.ProductID = order.ProductID
		resp.Behavior = order.Behavior
		resp.Price = order.Price
		resp.Amount = order.Amount
		resp.Status = order.Status
		resp.PaidAt = order.PaidAt
		resp.Product = structer.Product{}
		resp.Product.Id = order.Product.Id.Id
		resp.Product.Name = order.Product.Name
		resp.Product.Amount = order.Product.Amount
	}

	return
}
