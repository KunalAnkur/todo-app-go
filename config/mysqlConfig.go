package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func mysqlConnect() {
	dsn := "root:mysql@123@tcp(127.0.0.1:3306)/todostore?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	mysqlConnect()
	fmt.Println("MySql connection successfully")

	return db
}
