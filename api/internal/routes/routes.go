package routes

import (
	"go-test/internal/controllers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/logout", controllers.Logout)
	auth.Post("/register", controllers.Register)
}
