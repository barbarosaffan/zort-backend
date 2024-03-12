package router

import (
	"github.com/barbarosaffan/zort-backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// App Routes
	app.Get("/", handlers.HelloApp)

	app.Get("/ping", handlers.PingApp)

	// Zort Routes
	app.Get("/zorts", handlers.GetZorts)

	app.Get("/zorts/:id", handlers.GetZort)

	app.Post("/zorts", handlers.CreateZort)

	app.Put("/zorts/:id", handlers.UpdateZort)

	app.Delete("/zorts/:id", handlers.DeleteZort)
}
