package config

import (
	"errors"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"os"
)

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      os.Getenv("APP_NAME"),
		ErrorHandler: NewErrorHandler(),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		//
		code := fiber.StatusInternalServerError
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		var ve validator.ValidationErrors
		errorList := validation.GetError(err, ve)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": errorList,
			})
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
