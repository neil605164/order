package main

import (
	"embed"
	"order/app/global"
	"order/app/models"
	"order/internal/database"

	"gorm.io/gorm"
)

var f embed.FS

func init() {
	global.Start(f)
}

func main() {

	mysql := database.Instance()
	db, _ := mysql.DBConn()
	createTable(db,
		models.Order{},
		models.Product{},
	)

	println("資料表創建成功...")
}

func createTable(db *gorm.DB, tables ...interface{}) {
	for _, v := range tables {
		isExist := db.Migrator().HasTable(v)
		if isExist {
			if err := db.Migrator().DropTable(v); err != nil {
				println("drop table error : %v", err)
				panic(err)
			}
		}
		if err := db.Migrator().CreateTable(v); err != nil {
			println("create table error : %v", err)
			panic(err)
		}
	}
}
