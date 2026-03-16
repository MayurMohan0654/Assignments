package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *EnvConfig) {

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dns))

	if err != nil {
		panic("Failed to connect database")
	} else {
		println("connected to db")
	}

	DB = db
}
