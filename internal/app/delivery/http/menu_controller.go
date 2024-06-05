package http

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type MenuController struct {
	Log         *logrus.Logger
	Validator   *validator.Validate
	MenuUseCase usecase.InterfaceMenuUseCase
}

func NewMenuController(log *logrus.Logger, val *validator.Validate, mu usecase.InterfaceMenuUseCase) *MenuController {
	return &MenuController{
		Log:         log,
		Validator:   val,
		MenuUseCase: mu,
	}
}

func (c *MenuController) Explore(ctx *fiber.Ctx) error {
	return nil
}
