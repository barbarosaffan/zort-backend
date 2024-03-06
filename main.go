package main

import (
	"github.com/barbarosaffan/zort-backend/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.Init()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err = r.Run()

	if err != nil {
		log.Fatalf("r.Run() error: %s", err)
	}

}
