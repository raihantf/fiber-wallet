package controllers

import (
	user_request "fiber-wallet/app/dto/request"
	user_services "fiber-wallet/app/services"

	"github.com/gofiber/fiber/v2"
)

func Register(ctx *fiber.Ctx) error {
	data := new(user_request.UserRegisterRequest)

	if err := ctx.BodyParser(data); err != nil {
		return err
	}

	result, err := user_services.UserRegisterService(data)
	if err != nil {
		panic(err)
	}

	return ctx.JSON(&fiber.Map{
		"message": result,
	})
}
