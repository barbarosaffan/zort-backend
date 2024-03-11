package main

import (
	"fmt"

	"github.com/barbarosaffan/zort-backend/config"
	"github.com/barbarosaffan/zort-backend/database"
)

func main() {
	config.SetupEnv()
	database.Connect()

	fmt.Println("Server is running on port " + config.API_PORT)

}
