package orderb

import "order/app/global/errorcode"

// DeleteOrderById 透過 id 刪除訂單
func (b *business) DeleteOrderById(id string) (apiErr errorcode.Error) {

	// 刪除 db 資料
	if apiErr = b.order.DeleteOrderById(id); apiErr != nil {
		return
	}

	return
}
