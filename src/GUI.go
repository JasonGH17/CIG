package main

import "github.com/gofiber/fiber/v2"

func GUI() {
	app := fiber.New()
	app.Static("/", "./public")
	app.Get("/projects", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Send(getPrjsJson())
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})
	app.Get("/p", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})
	app.Listen(":8000")
}
