package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// type Middle struct {
// 	set *fiber.Map
// }

func Ping(c *fiber.Ctx) error {
	if c.Locals("email") != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": c.Locals("name")})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "not ok"})
}
