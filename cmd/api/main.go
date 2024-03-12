package main

import (
	"fmt"

	"github.com/barbarosaffan/zort-backend/config"
	"github.com/barbarosaffan/zort-backend/database"
	"github.com/barbarosaffan/zort-backend/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.SetupEnv()
	database.Connect()

	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":" + config.API_PORT)
	fmt.Println("Server is running on port " + config.API_PORT)
}
