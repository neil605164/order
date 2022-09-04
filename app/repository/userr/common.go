package userr

import (
	"order/app/global/errorcode"
	"order/app/global/structer"
	"order/app/models"
	"order/internal/database"
	"sync"
)

type Interface interface {
	UserList() (data []models.User, apiErr errorcode.Error)
	CreateUser(raw *structer.CreateReq) (apiErr errorcode.Error)
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
