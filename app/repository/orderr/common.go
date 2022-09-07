package orderr

import (
	"order/app/global/errorcode"
	"order/app/models"
	"order/internal/database"
	"sync"

	"github.com/shopspring/decimal"
)

type Interface interface {
	// 取 order list 資料
	GetOrderList() (orders []models.Order, apiErr errorcode.Error)
	// 建立訂單
	CreateOrder(order models.Order) (dbOrder models.Order, apiErr errorcode.Error)
	// 透過 id 取 order 資料
	GetOrderById(id string) (order models.Order, apiErr errorcode.Error)
	// 透過 id 刪除訂單
	DeleteOrderById(id string) (apiErr errorcode.Error)
	// 透過 id 更新訂單
	UpdateOrderById(id string, order map[string]interface{}) (apiErr errorcode.Error)
	// 透過行為與價格進行撮合
	GetOrderByBehaviorAndPrice(behavior string, price decimal.Decimal, productId uint64) (orders []models.Order, apiErr errorcode.Error)
}

var singleton *repo
var once sync.Once

type repo struct {
	DB database.IMySQL
}

// Instance 獲得單例對象
func Instance() Interface {
	once.Do(func() {
		singleton = &repo{
			DB: database.Instance(),
		}
	})
	return singleton
}
