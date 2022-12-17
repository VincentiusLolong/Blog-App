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

		name := sess.Get("logged_in")
		fmt.Println(name)
		str := fmt.Sprintf("%v", name)
		// token := c.Cookies("logged_in")
		// split := token.Split(string(token[:]), " ") // change split cookies

		data, claim := secure.ValidateToken(str)
		if claim != nil {
			return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{
				"status":  "ERROR",
				"message": claim.Error(),
			})
		}
		id := data["id"].(string)

		// c.Locals("token", split[1])
		// c.Locals("test", name)
		c.Locals("id", id)

		return c.Next()
	}
}
