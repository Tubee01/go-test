package main

import (
	"log"

	"go-test/internal/modules/config"
	"go-test/internal/modules/database"
	"go-test/internal/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config.Init()
	database.Init()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(config.AppConfig.APP_LISTEN))
}
