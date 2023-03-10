package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func ConnectDB() {
	dsn := "root:R^b2EV@!L@Io588Y@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(
		mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		log.Fatal("Cannot connect to the database")
	}

	Conn = db
	log.Fatal("Connect to the database......")
}
