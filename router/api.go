package router

import (
	"net/http"
	"order/app/handler/orderh"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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

}
