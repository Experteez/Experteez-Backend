package route

import (
	"Experteez-Backend/handler"
	"Experteez-Backend/handler/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	api := app.Group("/api")

	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.Post("/login", handler.UserLogin)
	users.Post("/talent/register", handler.UserTalentRegister)
	users.Get("/current", middleware.Auth, handler.GetCurrentUser)

	v1.Get("/projects", handler.ProjectHandlerGetAll)
}
