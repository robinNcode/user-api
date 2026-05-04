package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var username = "robin_ip_user"
var password = "dsnext79"
// Host only (no port); port is appended as :3306 in the DSN.
var host = "192.168.26.65"
var port = "3306"
var database = "user_service"

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("database connection failed: %w", err))
	}

	DB = db
}