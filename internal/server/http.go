package server

import (
	"context"
	"fmt"
	"net/http"
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/router"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Run HTTP 啟動 restful 服務
func Run() {
	helper.SetReportCallerSkip(3)

	defer func() {
		if err := recover(); err != nil {
			// 補上將err傳至telegram
			_ = helper.ErrorHandle(global.FatalLog, errorcode.Code.UnExpectedError, err)
			fmt.Println("[❌ Fatal❌ ] HTTP:", err)
		}
	}()

	_ = helper.ErrorHandle(global.SuccessLog, errorcode.Code.HTTPServerStart, "🔔 Run Http Service 🔔")

	// 本機開發需要顯示 Gin Log
	var r *gin.Engine
	if os.Getenv("ENV") == "local" {
		r = gin.New()
		r.Use(gin.Logger())
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// 自定義 Recovery
	r.Use(func(c *gin.Context) {
		defer helper.CatchError(c)
		c.Next()
	})

	// 載入router設定
	router.RouteProvider(r)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			_ = helper.ErrorHandle(global.SuccessLog, errorcode.Code.TCPPortDuplicate, err)
			fmt.Println("[❌ Fatal❌ ] Server 建立監聽連線失敗:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	_ = helper.ErrorHandle(global.SuccessLog, errorcode.Code.PrePareShutDownService, "🚦  收到訊號囉，等待其他連線完成，準備結束服務 🚦")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		_ = helper.ErrorHandle(global.SuccessLog, errorcode.Code.ServiceAlreadyShutdown, "🚦  收到關閉訊號，強制結束 🚦")
	}

}
