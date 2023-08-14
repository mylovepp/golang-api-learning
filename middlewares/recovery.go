package middlewares

import (
	"fmt"
	"log"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

func RecoveryMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			println("-----> RecoveryMiddleware called")
			err, ok := r.(error)

			if ok {
				log.Printf("Recovered from panic: %v\n", err)
			} else {
				log.Printf("Recovered from panic: %v\n", r)
			}
			printStackTrace()
			// You can choose to send a specific response to the client here
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
	}()
	return c.Next()
}

func printStackTrace() {
	stack := make([]byte, 8192)
	length := runtime.Stack(stack, false)
	fmt.Printf("Stack Trace:\n%s\n", stack[:length])
}
