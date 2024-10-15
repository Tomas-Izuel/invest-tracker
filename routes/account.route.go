package routes

import (
	"invest/models/dto"
	"invest/services"

	"github.com/gofiber/fiber/v2"
)

func AccountRoutes(app fiber.Router) {
	accountApi := app.Group("/accounts/:accountId")

	accountApi.Put("/", updateAccount)
	accountApi.Delete("/", deleteAccount)

	accountApi.Post("/add-investment", createInvestment)
}

func updateAccount(c *fiber.Ctx) error {
	accountId := c.Params("accountId")

	var accountDTO dto.UpdateAccountDTO

	if err := c.BodyParser(&accountDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	updatedAccount, err := services.UpdateAccount(c.Context(), accountId, &accountDTO)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Account updated successfully",
		"account": updatedAccount,
	})
}

func deleteAccount(c *fiber.Ctx) error {
	accountId := c.Params("accountId")

	err := services.DeleteAccount(c.Context(), accountId)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Account deleted successfully",
	})
}

func createInvestment(c *fiber.Ctx) error {
	accountId := c.Params("accountId")

	var investmentDTO dto.CreateInvestmentDTO

	if err := c.BodyParser(&investmentDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	createdInvestment, err := services.CreateInvestment(c.Context(), accountId, &investmentDTO)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":    "Investment created successfully",
		"investment": createdInvestment,
	})
}
