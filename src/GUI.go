package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type CLIPP struct {
	Op       int    `json:"op"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

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

	app.Post("/cli/project", func(c *fiber.Ctx) error {
		payload := CLIPP{}

		if err := c.BodyParser(&payload); err != nil {
			log.Printf("CLI Project API error: %v\n", err)
			return err
		}

		switch payload.Op {
		case 0:
			err := newProject(payload.Location, payload.Name)
			if err != nil {
				log.Printf("CLI Project API (new project) error: %v\n", err)
				return c.Send([]byte(err.Error()))
			}
			return c.SendStatus(201)
		default:
			return &CError{msg: "Invalid CLI Project API opcode"}
		}
	})

	app.Listen(":8000")
}
