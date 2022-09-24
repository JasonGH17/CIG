package main

import "github.com/gofiber/fiber/v2"

func GUI() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Send(getPrjsJson())
	})
	app.Listen(":8000")
}
