package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Get: Hello, World!")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Post: Hello, World!")
	})

	log.Fatal(app.Listen(":8080"))
}
