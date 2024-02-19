package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	user := "root"
	password := "root"
	ip := "localhost"
	port := "3306"
	dbname := "picobin"

	dsn := user + ":" + password + "@tcp(" + ip + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Database connection failure！ 数据库连接失败！")
	}
	DB = db
}
