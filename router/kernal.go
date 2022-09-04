package router

import (
	"order/app/global/helper"
	"order/app/middleware"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	// swagger embed files
	swaggerfiles "github.com/swaggo/files"
)

// RouteProvider 路由提供者
func RouteProvider(r *gin.Engine) {
	// 組合log基本資訊
	if helper.IsDeveloperEnv() {
		r.Use(middleware.WriteLog)
		// Swagger
		r.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	// api route
	LoadBackendRouter(r)
}
