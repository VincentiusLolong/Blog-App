package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// type Middle struct {
// 	set *fiber.Map
// }

func Ping(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "ok"})
}
