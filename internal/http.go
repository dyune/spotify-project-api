package internal

import (
	"github.com/gofiber/fiber/v2"
)

func StartApplication() *fiber.App {
	app := fiber.New()
	return app
}
