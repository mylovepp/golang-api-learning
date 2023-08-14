package routes

import (
	"log"
	"math/rand"
	"mylovepp/middlewares"
	"time"

	"github.com/goccy/go-json"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(middlewares.CorsMiddleware)
	app.Use(middlewares.RecoveryMiddleware)

	setupWebSocket(app)

	api := app.Group("/api")

	UserRoute(api)

	return app
}

func setupWebSocket(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/pricing", websocket.New(func(c *websocket.Conn) {
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		for {

			time.Sleep(time.Second)
			c.WriteJSON(fiber.Map{
				"code":      "RMF1",
				"assetType": "FUNDS",
				"assetName": "RMB Fixed Income Fund",
				"nav":       rand.Float64() * 100,
			})
		}
	}))

	app.Get("/ws/trading", websocket.New(func(c *websocket.Conn) {
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		var (
			msg []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			c.WriteJSON(fiber.Map{
				"code":   string(msg),
				"status": "SUCCESS",
			})
		}
	}))
}
