package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsMiddleware(c *fiber.Ctx) error {
	println("-----> CorsMiddleware called")
	allowHeaders := []string{
		fiber.HeaderOrigin,
		fiber.HeaderContentType,
		fiber.HeaderAccept,
		fiber.HeaderAuthorization,
		"x-api-key",
	}
	corsOrigin := cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5500",
		AllowHeaders:     strings.Join(allowHeaders, ","),
		AllowCredentials: true,
	})
	return corsOrigin(c)
}
