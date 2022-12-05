package middleware

import (
	"fiber-mongo-api/controllers/secure"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Auth() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("token")

		fmt.Println(token)

		// split := strings.Split(string(token[:]), " ")

		// if split[0] != "Bearer" {
		// 	return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{
		// 		"status":  "ERROR",
		// 		"message": "invalid authorization",
		// 	})
		// }
		data, claim := secure.ValidateToken(string(token))
		if claim != nil {
			return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{
				"status":  "ERROR",
				"message": claim.Error(),
			})
		}
		id := data["id"].(string)

		// c.Locals("token", split[1])
		c.Locals("id", id)

		return c.Next()
	}
}
