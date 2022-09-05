package productr

import (
	"order/app/global/errorcode"
	"order/app/models"
	"order/internal/database"
	"sync"
)

type Interface interface {
	// 透過產品 id 取產品資料
	GetProductById(id uint64) (product models.Product, apiErr errorcode.Error)
	// 檢查產品是否存在
	CheckProudctExistsById(id uint64) (isExist bool, apiErr errorcode.Error)
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
