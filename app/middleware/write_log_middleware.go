package middleware

import (
	"order/app/global/helper"

	"github.com/gin-gonic/gin"
)

// WriteLog 執行任何router前，都會紀錄一筆access.log
func WriteLog(c *gin.Context) {
	// 	寫access Log
	helper.Access(c)

	c.Next()
}
