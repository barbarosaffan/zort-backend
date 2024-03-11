package database

import (
	"log"

	"github.com/barbarosaffan/zort-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	dsn := "host=" + config.DB_HOST + " user=" + config.DB_USER + " password=" + config.DB_PASSWORD + " dbname=" + config.DB_NAME + " port=" + config.DB_PORT + " sslmode=disable TimeZone=Asia/Shanghai"

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("Database connected successfully.")

}
