package main

import (
	"Experteez-Backend/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	app := fiber.New()

	app.Listen(":3000")
}