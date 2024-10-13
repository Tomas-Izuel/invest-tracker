package routes

import (
	"invest/models/dto"
	"invest/services"

	"github.com/gofiber/fiber/v2"
)

func UserSignUp(c *fiber.Ctx) error {
	var userDTO dto.CreateUserDTO

	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := services.CreateUser(c.Context(), &userDTO); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
