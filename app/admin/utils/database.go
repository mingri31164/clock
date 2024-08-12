package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 开启连接池
func InitDB() *gorm.DB {
	host := "116.205.189.126"
	port := "3306"
	database := "clock-go"
	username := "clock"
	password := "clock1234"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err: " + err.Error())
	}
	//自动创建数据表
	//db.AutoMigrate(&model.Test{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
