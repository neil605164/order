package orderb

import "order/app/global/errorcode"

func (b *business) UpdateOrderById(id string) (apiErr errorcode.Error) {

	//
	order := make(map[string]interface{})

	// 更新訂單
	b.order.UpdateOrderById(id, order)
	return
}
