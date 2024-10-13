package routes

import (
	"invest/models/dto"
	"invest/services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(app fiber.Router) {
	usersApi := app.Group("/users")

	usersApi.Post("/signup", userSignUp)
}

func userSignUp(c *fiber.Ctx) error {
	var userDTO dto.CreateUserDTO

	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	createdUser, err := services.CreateUser(c.Context(), &userDTO)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    createdUser,
	})
}
