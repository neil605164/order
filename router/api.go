package router

import (
	"net/http"
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/global/structer"
	"order/app/handler/orderh"
	"order/app/models"
	"order/internal/cache"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	// gin-swagger middleware
	// swagger embed files
)

// LoadBackendRouter 路由控制
func LoadBackendRouter(r *gin.Engine) {

	api := r.Group("/api/v1")

	// K8S Health Check
	api.GET("/healthz", func(c *gin.Context) {
		data := map[string]string{
			"service": os.Getenv("PROJECT_NAME"),
			"time":    time.Now().Format("2006-01-02 15:04:05 -07:00"),
		}

		c.JSON(http.StatusOK, data)
	})

	order := api.Group("/order")
	{
		order.GET("", orderh.OrderList)
		order.GET("/:id", orderh.OrderDetail)
		order.POST("", orderh.CreateOrder)
		// order.PUT("/:id", orderh.UpdateOrder)
		order.DELETE("/:id", orderh.DeleteOrder)
	}

	api.GET("", func(c *gin.Context) {

		order := models.Order{
			ProductID: 1,
			OrderNo:   "qwerasdf1234",
			Behavior:  "sell",
			Price:     decimal.NewFromInt(5),
			Amount:    10,
			Status:    "unpaid",
		}

		// 處理丟入 queue 資料
		byteData, err := jsoniter.Marshal(order)
		if err != nil {
			apiErr := helper.ErrorHandle(global.WarnLog, errorcode.Code.JSONMarshalError, nil)
			c.JSON(http.StatusOK, helper.Fail(apiErr))
			return
		}
		queueData := structer.RedisLPushFormat{
			Type: global.OrderQueue,
			Data: order,
		}

		byteData, err = jsoniter.Marshal(queueData)
		if err != nil {
			apiErr := helper.ErrorHandle(global.WarnLog, errorcode.Code.JSONMarshalError, nil)
			c.JSON(http.StatusOK, helper.Fail(apiErr))
			return
		}

		cache := cache.Instance()
		if err := cache.LPush(global.RedisQueueChannel, byteData); err != nil {
			apiErr := helper.ErrorHandle(global.WarnLog, errorcode.Code.ProductNotExist, nil, string(byteData))
			c.JSON(http.StatusOK, helper.Fail(apiErr))
			return
		}

		c.JSON(http.StatusOK, "success")
	})

}
