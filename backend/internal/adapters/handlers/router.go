package handlers

import (
	"GoNext/base/internal/core/ports"
	"GoNext/base/internal/core/services"
	"GoNext/base/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app *fiber.App
	authHandler *AuthHandler
	userHandler *UserHandler
}

func NewRouter(app *fiber.App, userRepo ports.UserRepository, jwtSecret string) *Router {

	authService := services.NewAuthService(userRepo, jwtSecret)
	userService := services.NewUserService(userRepo)
	
	authHandler := NewAuthHandler(authService, userService)
	userHandler := NewUserHandler(userService)

	return &Router{app: app, authHandler: authHandler, userHandler: userHandler}
}

func (r *Router) SetupPublicRoutes() {
	r.app.Post("/api/auth/register", r.authHandler.Register)
	r.app.Post("/api/auth/login", r.authHandler.Login)
	r.app.Post("/api/auth/logout", r.authHandler.Logout)
	r.app.Get("/api/auth/status", r.authHandler.Status)
}

func (r *Router) SetupProtectedRoutes() {
	api := r.app.Group("/api", middleware.JWTAuthentication(r.authHandler.authService))
	api.Get("/users/me", r.userHandler.GetCurrentUser)
	api.Get("/users", r.userHandler.GetByEmail)
}



