package internal

import "github.com/gofiber/fiber/v2"

func TestRoutes(app *fiber.App) {
	app.Get("/ping", ping)
}

func ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}
