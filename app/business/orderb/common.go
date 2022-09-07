package orderb

import (
	"order/app/global/errorcode"
	"order/app/global/structer"
	"order/app/repository/orderr"
	"order/app/repository/productr"
	"order/internal/cache"
	"sync"
)

var singleton *business
var once sync.Once

type IOrder interface {
	// 建立訂單
	CreateOrder(raw structer.CreateOrderReq) (apiErr errorcode.Error)
	// 取 order list
	GetOrderList() (resp []structer.OrderList, apiErr errorcode.Error)
	// 透過 id 取 order 資料
	GetOrderById(id string) (resp *structer.OrderList, apiErr errorcode.Error)
	// 透過 id 刪除訂單
	DeleteOrderById(id string) (apiErr errorcode.Error)
	// 撮合訂單
	MatchOrder(byteData []byte) (apiErr errorcode.Error)
}

type business struct {
	product productr.Interface
	order   orderr.Interface
	cache   cache.IRedis
}

func Instance() IOrder {
	once.Do(func() {
		singleton = &business{
			product: productr.Instance(),
			order:   orderr.Instance(),
			cache:   cache.Instance(),
		}
	})
	return singleton
}
