package handlers

import "github.com/gofiber/fiber/v2"

func PingApp(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func HelloApp(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello!",
	})
}
