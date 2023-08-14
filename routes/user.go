package routes

import (
	"mylovepp/controllers"
	"mylovepp/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(route fiber.Router) {
	apiV1 := route.Group("/v1/users")
	apiV1.Use(middlewares.VerifyApiKeyMiddleware)

	userController := controllers.UserController{}

	apiV1.Get("", userController.GetUsers)
	apiV1.Get("/:id", userController.GetUser)
	apiV1.Post("", userController.AddUser)
	apiV1.Put("/:id", userController.UpdateUser)
	apiV1.Delete("/:id", userController.DeleteUser)
	apiV1.Patch("", userController.AddUserWithAccounts)

	apiV2 := route.Group("/v2/users")
	apiV2.Get("", userController.GetUsers)

}
