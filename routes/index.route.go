package routes

import "github.com/gofiber/fiber/v2"

func handler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func RouteHandler(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/test", handler)

	UsersRoutes(api)
}
