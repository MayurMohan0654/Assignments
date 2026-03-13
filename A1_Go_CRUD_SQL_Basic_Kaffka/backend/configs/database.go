package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	db, err := gorm.Open(mysql.Open("root:xgok@tcp(127.0.0.1:3306)/a1?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		panic("Failed to connect database")
	}else{
		println("connected to db");
	}

	DB = db
}