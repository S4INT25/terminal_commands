package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"terminal_commands/utils"
)

func Authentication(c *fiber.Ctx) error {

	token := c.Get("Authorization")

	if token != "" {

		claims, err := utils.ValidateToken(token)

		if err != nil {

			return c.SendString(fmt.Sprintf("Invalid token %v", err))
		}

		id := fmt.Sprintf("%v", claims["id"])
		email := fmt.Sprintf("%v", claims["email"])

		c.Locals("id", id)
		c.Locals("email", email)

	} else {
		return c.SendString("authorization token not found")
	}

	return c.Next()
}
