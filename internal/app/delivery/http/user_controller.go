package http

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/usecase"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log         *logrus.Logger
	Validator   *validator.Validate
	UserUseCase usecase.InterfaceUserUseCase
}

func NewUserController(log *logrus.Logger, val *validator.Validate, uu usecase.InterfaceUserUseCase) *UserController {
	return &UserController{
		Log:         log,
		Validator:   val,
		UserUseCase: uu,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RequestUserRegister)

	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.Validator.Struct(request); err != nil {
		fmt.Println(request)
		return err
	}

	// usecase
	res, err := c.UserUseCase.Create(request)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Success{
		Message: "Success Register User",
		Data:    res,
		Status: response.Status{
			Code:    fiber.StatusOK,
			Message: "OK",
		},
	})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.RequestUserLogin)

	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.Validator.Struct(request); err != nil {
		return err
	}

	res, err := c.UserUseCase.Login(request)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Success{
		Message: "Success Login User",
		Data:    res,
		Status: response.Status{
			Code:    fiber.StatusOK,
			Message: "OK",
		},
	})
}

func (c *UserController) ForgotPassword(ctx *fiber.Ctx) error {
	return ctx.JSON("")
}
