package database

import (
	"fmt"
	"log"
	"order/app/global"
	"order/app/global/errorcode"
	"order/app/global/helper"
	"order/app/models"
	"os"
	"sync"
	"time"

	"gorm.io/plugin/dbresolver"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IMySQL interface {
	DBConn() (*gorm.DB, errorcode.Error)
	DBPing()
	CheckTable()
}

// MySQL DB連線資料
type MySQL struct {
	host     string `tag:"DB host"`
	username string `tag:"DB username"`
	password string `tag:"DB password"`
	database string `tag:"DB database"`
}

var singleton *MySQL
var once sync.Once
var dbCon *gorm.DB

func Instance() IMySQL {
	once.Do(func() {
		singleton = &MySQL{}
	})
	return singleton
}

func (m *MySQL) DBConn() (*gorm.DB, errorcode.Error) {

	if dbCon != nil {
		return dbCon, nil
	}

	var err error

	dsnMaster := composeString(global.DBMaster)
	dsnSlave := composeString(global.DBSlaver)

	// 連接gorm
	dbCon, err = gorm.Open(mysql.Open(dsnMaster), &gorm.Config{})
	if err != nil {
		apiErr := helper.ErrorHandle(global.FatalLog, errorcode.Code.DBConnectError, err.Error())
		return nil, apiErr
	}

	_ = dbCon.Use(
		dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dsnMaster)},
			Replicas: []gorm.Dialector{mysql.Open(dsnSlave)},
			Policy:   dbresolver.RandomPolicy{}}).
			// 空閒連線 timeout 時間
			SetConnMaxIdleTime(15 * time.Second).
			// 空閒連線 timeout 時間
			SetConnMaxLifetime(15 * time.Second).
			// 限制最大閒置連線數
			SetMaxIdleConns(100).
			// 限制最大開啟的連線數
			SetMaxOpenConns(2000),
	)

	sqlDB, _ := dbCon.DB()

	// 限制最大閒置連線數
	sqlDB.SetMaxIdleConns(100)
	// 限制最大開啟的連線數
	sqlDB.SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	sqlDB.SetConnMaxLifetime(15 * time.Second)

	if global.Config.DB.Debug {
		dbCon = dbCon.Debug()
	}

	return dbCon, nil
}

// DBPing 檢查DB是否啟動
func (m *MySQL) DBPing() {
	// 檢查 master db
	dbCon, apiErr := m.DBConn()
	if apiErr != nil {
		log.Fatalf("🔔🔔🔔 MASTER DB CONNECT ERROR: %v 🔔🔔🔔", global.Config.DBMaster.Host)
	}

	masterDB, err := dbCon.DB()
	if err != nil {
		log.Fatalf("🔔🔔🔔 CONNECT MASTER DB ERROR: %v 🔔🔔🔔", err.Error())
	}

	err = masterDB.Ping()
	if err != nil {
		log.Fatalf("🔔🔔🔔 PING MASTER DB ERROR: %v 🔔🔔🔔", err.Error())
	}

}

// CheckTable 啟動main.go服務時，直接檢查所有 DB 的 Table 是否已經存在
func (m *MySQL) CheckTable() {
	db, apiErr := m.DBConn()
	if apiErr != nil {
		log.Fatalf("🔔🔔🔔 MASTER DB CONNECT ERROR: %v 🔔🔔🔔", global.Config.DBMaster.Host)
	}

	// 會自動建置 DB Table
	err := db.Set("gorm:table_options", "comment '訂單'").AutoMigrate(&models.Order{})
	if err != nil {
		panic(err.Error())
	}

	err = db.Set("gorm:table_options", "comment '產品'").AutoMigrate(&models.Product{})
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		_ = helper.ErrorHandle(global.FatalLog, errorcode.Code.DBTableNotExist, fmt.Sprintf("❌ 設置DB錯誤： %v ❌", err.Error()))
		log.Fatalf("🔔🔔🔔 MIGRATE MASTER TABLE ERROR: %v 🔔🔔🔔", err.Error())
	}

}

// composeString 組合DB連線前的字串資料
func composeString(mode string) string {
	db := MySQL{}

	switch mode {
	case global.DBMaster:
		db.host = getHost(true)
		db.username = getUser(true)
		db.password = getPass(true)
		db.database = getDBName()
	case global.DBSlaver:
		db.host = getHost(false)
		db.username = getUser(false)
		db.password = getPass(false)
		db.database = getDBName()
	}

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?timeout=5s&charset=utf8mb4&parseTime=True&loc=Local", db.username, db.password, db.host, db.database)
}

// getHost 取 DB host
func getHost(master bool) string {

	switch master {
	case true:
		if value, ok := os.LookupEnv("MHOST"); ok {
			return value
		}
		return global.Config.DBMaster.Host
	default:
		if value, ok := os.LookupEnv("SHOST"); ok {
			return value
		}
		return global.Config.DBMaster.Host
	}
}

// getUser 取 DB User
func getUser(master bool) string {

	switch master {
	case true:
		if value, ok := os.LookupEnv("MUSER"); ok {
			return value
		}
		return global.Config.DBMaster.Username
	default:
		if value, ok := os.LookupEnv("SUSER"); ok {
			return value
		}
		return global.Config.DBMaster.Username
	}
}

// getPass 取 DB Pass
func getPass(master bool) string {

	switch master {
	case true:
		if value, ok := os.LookupEnv("MPASS"); ok {
			return value
		}
		return global.Config.DBMaster.Password
	default:
		if value, ok := os.LookupEnv("SPASS"); ok {
			return value
		}
		return global.Config.DBMaster.Password
	}
}

// getDBName 取 DB Name
func getDBName() string {

	if value, ok := os.LookupEnv("DB_NAME"); ok {
		return value
	}
	return global.Config.DBMaster.Database
}
