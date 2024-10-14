package routes

import (
	"invest/models/dto"
	"invest/services"

	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(app fiber.Router) {
	usersApi := app.Group("/users")

	usersApi.Post("/signup", userSignUp)
	usersApi.Post("/signin", userSignIn)

	existUserApi := usersApi.Group("/:id")

	existUserApi.Post("/add-account", addAccount)
	existUserApi.Get("/accounts", getAccounts)
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

func userSignIn(c *fiber.Ctx) error {
	var userDTO dto.SignInDTO

	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	user, err := services.SignIn(c.Context(), &userDTO)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User signed in successfully",
		"user":    user,
	})
}

func addAccount(c *fiber.Ctx) error {
	userId := c.Params("id")

	var accountDTO dto.CreateAccountDTO

	if err := c.BodyParser(&accountDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	createdAccount, err := services.CreateAccount(c.Context(), userId, &accountDTO)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Account created successfully",
		"account": createdAccount,
	})
}

func getAccounts(c *fiber.Ctx) error {
	userId := c.Params("id")

	accounts, err := services.GetAccountByUserID(c.Context(), userId)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(accounts)
}
