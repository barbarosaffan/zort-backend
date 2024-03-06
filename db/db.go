package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	opt, err := pg.ParseURL(os.Getenv("DB_URL"))

	if err != nil {
		log.Fatalf("Error parsing DB_URL: %s", err)
	}

	pg.Connect(opt)
}
