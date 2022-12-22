package middleware

import (
	"fiber-mongo-api/controllers"
	"fiber-mongo-api/controllers/secure"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Auth() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := controllers.Store.Get(c)
		if err != nil {
			panic(err)
		}

		name := sess.Get("logged_in") // Add Err
		fmt.Println(name)
		str := fmt.Sprintf("%v", name)

		data, claim := secure.ValidateToken(str)
		if claim != nil {
			return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{
				"status":  "ERROR",
				"message": claim.Error(),
			})
		}

		id := data["id"].(string)
		if id == "000000000000000000000000" {
			return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{
				"status":  "ERROR",
				"message": "Error : Cant Find User_id",
			})
		}

		c.Locals("id", id)

		return c.Next()
	}
}
