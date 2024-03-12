package database

import (
	"log"

	"github.com/barbarosaffan/zort-backend/config"
	"github.com/barbarosaffan/zort-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	dsn := "host=" + config.DB_HOST + " user=" + config.DB_USER + " password=" + config.DB_PASSWORD + " dbname=" + config.DB_NAME + " port=" + config.DB_PORT + " sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db

	log.Println("Database connected successfully.")

	err = DB.AutoMigrate(&models.Zort{})

	if err != nil {
		panic(err)
	}

	return DB
}
