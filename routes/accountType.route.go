package routes

import (
	"invest/models/dto"
	"invest/services"

	"github.com/gofiber/fiber/v2"
)

func AccountTypeRoutes(app fiber.Router) {
	accountTypeApi := app.Group("/account-types")

	accountTypeApi.Post("/create", createAccountType)
	accountTypeApi.Get("/", getAllAccountTypes)
	accountTypeApi.Put("/:id", updateAccountType)
}

func createAccountType(c *fiber.Ctx) error {
	var accountTypeDTO dto.CreateAccountTypeDTO

	if err := c.BodyParser(&accountTypeDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	createdAccountType, err := services.CreateAccountType(c.Context(), &accountTypeDTO)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Account type created successfully",
		"account": fiber.Map{
			"id":   createdAccountType.InsertedID,
			"name": accountTypeDTO.Name,
		},
	})
}

func getAllAccountTypes(c *fiber.Ctx) error {
	accountTypes, err := services.GetAllAccountTypes(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"accountTypes": accountTypes,
	})
}

func updateAccountType(c *fiber.Ctx) error {
	var accountTypeDTO dto.CreateAccountTypeDTO
	accountTypeID := c.Params("id")

	if err := c.BodyParser(&accountTypeDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	_, err := services.UpdateAccountType(c.Context(), accountTypeID, &accountTypeDTO)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Account type updated successfully",
		"account": fiber.Map{
			"id":   accountTypeID,
			"name": accountTypeDTO.Name,
		},
	})
}
