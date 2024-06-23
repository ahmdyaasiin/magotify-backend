package config

import (
	"errors"
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/response"
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
		if errors.As(err, &ve) {
			errorList := validation.GetError(err, ve)
			if err != nil {
				fmt.Println(err)
				return ctx.Status(fiber.StatusBadRequest).JSON(response.ValidationError{
					Message: "Validation Error",
					Errors:  errorList,
					Status: response.Status{
						Code:    fiber.StatusBadRequest,
						Message: "Validation Error",
					},
				})
			}
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
