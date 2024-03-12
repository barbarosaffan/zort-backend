package handlers

import (
	"strconv"

	"github.com/barbarosaffan/zort-backend/models"
	"github.com/barbarosaffan/zort-backend/repository"
	"github.com/gofiber/fiber/v2"
)

func GetZorts(c *fiber.Ctx) error {
	zorts, err := repository.FindAllZorts()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "An error occurred while fetching zorts.",
		})
	}

	return c.JSON(zorts)
}

func GetZort(c *fiber.Ctx) error {
	id := c.Params("id")

	uintId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID.",
		})
	}

	zort, err := repository.FindZortByID(uintId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Zort not found.",
		})
	}

	return c.JSON(zort)
}

func CreateZort(c *fiber.Ctx) error {
	zort := new(models.Zort)

	if err := c.BodyParser(zort); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request.",
		})
	}

	if err := repository.CreateZort(zort); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "An error occurred while creating zort.",
		})
	}

	return c.JSON(zort)
}

func UpdateZort(c *fiber.Ctx) error {
	id := c.Params("id")

	uintId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID.",
		})
	}

	zort, err := repository.FindZortByID(uintId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Zort not found.",
		})
	}

	if err := c.BodyParser(&zort); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request.",
		})
	}

	if err := repository.UpdateZort(&zort); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "An error occurred while updating zort.",
		})
	}

	return c.JSON(zort)
}

func DeleteZort(c *fiber.Ctx) error {
	id := c.Params("id")

	uintId, err := strconv.ParseUint(id, 10, 64)

	zort, err := repository.FindZortByID(uintId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Zort not found.",
		})
	}

	if err := repository.DeleteZort(&zort); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "An error occurred while deleting zort.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Zort deleted successfully.",
	})
}
