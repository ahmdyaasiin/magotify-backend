package http

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http/middleware"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/usecase"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/response"
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

	auth := middleware.GetUser(ctx)

	res, err := c.MenuUseCase.Explore(auth)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Success{
		Message: "Success get Explore",
		Data:    res,
		Status: response.Status{
			Code:    fiber.StatusOK,
			Message: "OK",
		},
	})
}

func (c *MenuController) Shop(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	res, err := c.MenuUseCase.Shop(auth)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Success{
		Message: "Success get Shop",
		Data:    res,
		Status: response.Status{
			Code:    fiber.StatusOK,
			Message: "OK",
		},
	})
}
