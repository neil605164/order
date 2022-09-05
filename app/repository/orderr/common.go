package orderr

import (
	"order/app/global/errorcode"
	"order/internal/database"
	"sync"
)

type Interface interface {
	// 建立訂單
	CreateOrder(order map[string]interface{}) (apiErr errorcode.Error)
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
