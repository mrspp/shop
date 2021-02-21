package main

import (
	"backend/migrations"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// config.GetDbConnection()
	migrations.Migrate()
	// helpers.GetCategory("https://shopee.vn/api/v4/official_shop/get_categories?tab_type=1")
	app := fiber.New()
	app.Static("/", "./public")

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3001")
}
