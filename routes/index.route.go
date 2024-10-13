package routes

import "github.com/gofiber/fiber/v2"

func handler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func RouteHandler(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/list", handler)
	v1.Get("/user", handler)

	v2 := api.Group("/v2")
	v2.Get("/list", handler)
	v2.Get("/user", handler)
}
