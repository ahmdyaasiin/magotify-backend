package http

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http/middleware"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/usecase"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CartController struct {
	Log         *logrus.Logger
	Validator   *validator.Validate
	CartUseCase usecase.InterfaceCartUseCase
}

func NewCartController(log *logrus.Logger, val *validator.Validate, cu usecase.InterfaceCartUseCase) *CartController {
	return &CartController{
		Log:         log,
		Validator:   val,
		CartUseCase: cu,
	}
}

func (c *CartController) GetCart(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	res, err := c.CartUseCase.GetCart(auth)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Success{
		Message: "Success get Cart",
		Data:    res,
		Status: response.Status{
			Code:    fiber.StatusOK,
			Message: "OK",
		},
	})
}

func (c *CartController) AddToCart(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.RequestAddCart)

	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.Validator.Struct(request); err != nil {
		return err
	}

	// usecase
	res, err := c.CartUseCase.AddCart(auth, request)
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
