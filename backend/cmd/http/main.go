package main

import (
	"log"
	"os"

	"GoNext/base/internal/adapters/handlers"
	"GoNext/base/internal/adapters/repositories"
	"GoNext/base/internal/core/services"
	"GoNext/base/pkg/database"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	entClient := database.NewEntClient(dbHost, dbPort, dbUser, dbPassword, dbName)
	defer entClient.Close()

	userRepo := repositories.NewUserRepository(entClient)

	userService := services.NewUserService(userRepo)

	userHandler := handlers.NewUserHandler(userService)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/users/create", userHandler.CreateUser)
	app.Post("/users/id", userHandler.GetById)
	app.Post("/users/email", userHandler.GetByEmail)

	log.Fatal(app.Listen(":8080"))
}
