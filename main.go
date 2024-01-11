package main

import (
	"Experteez-Backend/database"
	"Experteez-Backend/database/migrations"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	migrations.RunMigrations()

	app := fiber.New()

	// route.SetupRoutes(app)

	app.Listen(":8080")
}