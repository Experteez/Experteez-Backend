package handler

import (
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(c *fiber.Ctx) error {
	return c.SendString("Get all users")
}