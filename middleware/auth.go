package middleware

import (
	"fiber-mongo-api/controllers/secure"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Request().Header.Peek("Authorization")

		split := strings.Split(string(token[:]), " ")

		if split[0] != "Bearer" {
			return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{
				"status":  "ERROR",
				"message": "invalid authorization",
			})
		}
		claim := secure.ValidateToken(string(split[1]))
		if claim != nil {
			return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{
				"status":  "ERROR",
				"message": claim.Error(),
			})
		}
		// 	email := claim["email"].(string)
		// // id := userClaims["id"].(string)

		c.Locals("token", split[1])
		// c.Locals("user_email", email)
		return c.Next()
	}
}
