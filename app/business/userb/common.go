package userb

import (
	"order/app/global/errorcode"
	"order/app/global/structer"
	"order/app/repository/userr"
	"sync"
)

var singleton *business
var once sync.Once

type IUser interface {
	UserList() ([]structer.UserList, errorcode.Error)
	CreateUser(raw *structer.CreateReq) (apiErr errorcode.Error)
}

type business struct {
	user userr.Interface
}

func Instance() IUser {
	once.Do(func() {
		singleton = &business{
			user: userr.Instance(),
		}
	})
	return singleton
}
