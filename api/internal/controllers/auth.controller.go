package controllers

import (
	"go-test/internal/modules"
	"go-test/internal/services"

	"github.com/gofiber/fiber/v3"
)

func Login(ctx fiber.Ctx) error {
	data := new(services.LoginData)
	if err := ctx.Bind().Body(data); err != nil {
		return err
	}

	user, err := services.Login(data)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to login"})
	}

	_, err = modules.SetSession(ctx, "user", user)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to login"})
	}

	return ctx.SendStatus(200)
}

func Register(ctx fiber.Ctx) error {
	data := new(services.RegisterData)
	if err := ctx.Bind().Body(data); err != nil {
		return err
	}

	err := services.Register(data)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to register"})
	}

	return ctx.SendStatus(200)
}

func Logout(ctx fiber.Ctx) error {
	err := modules.DestroySession(ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to logout"})
	}

	return ctx.SendStatus(200)
}
