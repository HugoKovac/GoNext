package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"GoNext/base/internal/adapters/handlers"
	"GoNext/base/internal/adapters/repositories"
	"GoNext/base/pkg/database"
	"GoNext/base/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

func main() {
	config := config.LoadConfig()

	entClient := database.NewEntClient(config)
	defer entClient.Close()

	userRepo := repositories.NewUserRepository(entClient)

	
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.Cors.AllowOrigins,
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowCredentials: true,
	}))
	
	router := handlers.NewRouter(app, userRepo, config)
	router.SetupPublicRoutes()
	router.SetupProtectedRoutes()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-shutdown
		log.Println("Shutting down...")
		app.Shutdown()
		os.Exit(0)
	}()

	log.Fatal(app.Listen(":8080"))
}
