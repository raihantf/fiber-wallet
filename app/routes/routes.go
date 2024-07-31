package routes

import (
	user_controller "fiber-wallet/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Post("/register", user_controller.Register)
	app.Post("/login", user_controller.Login)
}
