package main

import (
	"embed"
	"order/app/global"
	"order/app/global/helper"
	"order/internal/cache"
	"order/internal/database"
	"order/internal/entry"

	_ "order/docs"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

//go:embed env/*
var f embed.FS

// 初始化動作
func init() {

	// 載入環境設定，所有動作須在該func後執行
	global.Start(f)

	// 設定 log 格式
	helper.SetFormatter(&logrus.JSONFormatter{})

	// 檢查 DB 機器服務
	database.Instance().DBPing()

	// 自動建置 DB + Table
	// if helper.IsLocal() {
	// 	database.Instance().CheckTable()
	// }

	// 檢查 Redis 機器服務
	_ = cache.Instance().Ping()

	// 設定程式碼 timezone
	// os.Setenv("TZ", "America/Puerto_Rico")
}

// @title Order Demo
// @version 1.0
// @description     This is a sample server celler server.
// @host      localhost:9999
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
func main() {
	entry.Run()
}
