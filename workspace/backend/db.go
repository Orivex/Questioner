package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:root@tcp(127.0.0.1:3306)/questionerdb?charset=utf8mb4&parseTime=True&loc=Local"

func initDatabase() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&User{}, &Question{})
}
