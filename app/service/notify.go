package service

import (
	"fmt"
	"order/app/global"
	"order/internal/cache"
	"runtime"
	"sync"
)

type notify struct {
	redisRepo cache.IRedis
}

var singleton *notify
var once sync.Once

type INotifier interface {
	HandleServiceError(org, url string, header map[string]string, params, resp interface{})
}

// NewNotify Implement NewNotify
func NewNotify() INotifier {
	once.Do(func() {
		singleton = &notify{
			redisRepo: cache.Instance(),
		}
	})
	return singleton
}

// SendMessage 發送訊息
func (n *notify) sendMessage(org string, content content) {

	// 取發送訊息識別證
	funcName, _, line, _ := runtime.Caller(2)
	f := runtime.FuncForPC(funcName)

	key := fmt.Sprintf(global.TeamPlus+":%v-%d", f.Name(), line)

	// Redis 撈資料，存在表示寄發過，直接忽略
	data, _ := n.redisRepo.Get(key)

	if data != "" {
		return
	}

	// 發送訊息
	// helper.TeamPlus(org, content)

	// 寫入 redis
	_ = n.redisRepo.Set(key, "true", global.RedisSendMessageExpire)
}

// handleServiceError service 錯誤處理
func (n *notify) HandleServiceError(org, url string, header map[string]string, params, resp interface{}) {
	// Teamplus 組訊息內容
	msg := content{
		Req: request{
			URL:    url,
			Header: header,
			Param:  params,
		},
		Resp: resp,
	}

	n.sendMessage(org, msg)
}
