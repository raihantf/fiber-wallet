package main

import (
	"fiber-wallet/app/database"
	"fiber-wallet/app/error"
	"fiber-wallet/app/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func InitFiber() *fiber.App {
	config := fiber.Config{
		ErrorHandler: error.CustomError,
	}
	return fiber.New(config)
}

func main() {
	godotenv.Load(".env")
	database.DBInit()

	app := InitFiber()
	app.Use(logger.New())

	routes.RouteInit(app)

	port := os.Getenv("PORT")

	appErr := app.Listen(":" + port)

	if appErr != nil {
		fmt.Println("Server Init Error!")
		os.Exit(1)
	}
}
