package redisqueue

import (
	"context"
	"errors"
	"fmt"
	"order/app/business/orderb"
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/global/structer"
	"order/internal/bootstrap"
	"order/internal/cache"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func Run(ctx context.Context) {
	timeout := 60 * time.Second

	defer func() {
		if err := recover(); err != nil {
			_ = helper.ErrorHandle(global.FatalLog, errorcode.Code.RedisSubscribeFail, fmt.Sprintf("%v", err))
		}
	}()

	_ = helper.ErrorHandle(global.SuccessLog, errorcode.Code.RedisSubscribeStart, "🔔 redis subscribe success start 🔔")

	sub := structer.RedisLPushFormat{}
	var c chan string
	for {

		select {
		case <-bootstrap.GracefulDown():
		case msg := <-c:
			if err := jsoniter.Unmarshal([]byte(msg), &sub); err != nil {
				_ = helper.ErrorHandle(global.WarnLog, errorcode.Code.JSONUnMarshalError, err.Error())
				return
			}

			switch sub.Type {
			case global.OrderQueue:
				defer func() {
					if err := recover(); err != nil {
						_ = helper.ErrorHandle(global.WarnLog, errorcode.Code.RedisOrderQueueUnexpectFail, err)
					}
				}()

				byteData, err := jsoniter.Marshal(sub.Data)
				if err != nil {
					_ = helper.ErrorHandle(global.WarnLog, errorcode.Code.JSONMarshalError, err)
					return
				}

				bus := orderb.Instance()
				_ = bus.MatchOrder(byteData)

			default:
				func() {
					defer func() {
						if err := recover(); err != nil {
							_ = helper.ErrorHandle(global.WarnLog, errorcode.Code.RedisQueueTypeNotRegister, nil)
						}
					}()
				}()

			}

		default:

			c = make(chan string, 1)
			cache := cache.Instance()
			subscriber := cache.BRPop(global.RedisQueueChannel, timeout)
			if subscriber.Err() != nil {
				if errors.Is(subscriber.Err(), fmt.Errorf("redis: nil")) {
					_ = helper.ErrorHandle(global.WarnLog, errorcode.Code.RedisSubscribeFail, subscriber.Err())
				}
			}

			if subscriber != nil {
				message := subscriber.Val()
				if len(message) > 0 {
					c <- message[1]
				}
			}
		}
	}
}
