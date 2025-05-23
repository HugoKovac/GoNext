package main

import (
	"log"
	"os"

	"GoNext/base/internal/adapters/handlers"
	"GoNext/base/internal/adapters/repositories"
	"GoNext/base/internal/core/services"
	"GoNext/base/internal/middleware"
	"GoNext/base/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_jwt_secret" // Not recommended for production
		log.Println("Warning: Using default JWT secret. Set JWT_SECRET environment variable for security.")
	}

	entClient := database.NewEntClient(dbHost, dbPort, dbUser, dbPassword, dbName)
	defer entClient.Close()

	userRepo := repositories.NewUserRepository(entClient)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, jwtSecret)

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService, userService)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://gonext.com",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowCredentials: true,
	}))

	app.Post("/api/auth/register", authHandler.Register)
	app.Post("/api/auth/login", authHandler.Login)
	app.Post("/api/auth/logout", authHandler.Logout)
	app.Get("/api/auth/status", authHandler.Status)

	// Protected routes
	api := app.Group("/api", middleware.JWTAuthentication(authService))
	api.Get("/users/me", userHandler.GetCurrentUser)
	api.Get("/users", userHandler.GetByEmail)

	log.Fatal(app.Listen(":8080"))
}
