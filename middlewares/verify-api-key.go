package middlewares

import "github.com/gofiber/fiber/v2"

func VerifyApiKeyMiddleware(c *fiber.Ctx) error {
	println("-----> VerifyApiKeyMiddleware called")
	apiKey := c.Request().Header.Peek("x-api-key")
	if apiKey == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	println("-----> apiKey: ", string(apiKey))

	return c.Next()
}
