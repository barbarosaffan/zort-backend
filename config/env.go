package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	API_PORT    string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_HOST     string
)

func SetupEnv() {
	err := godotenv.Load(".env")

	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env file")
	}

	API_PORT = os.Getenv("API_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	DB_HOST = os.Getenv("DB_HOST")

	log.Println("Environment loaded successfully.")
}
