package http

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http/middleware"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/usecase"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TransactionController struct {
	Log                *logrus.Logger
	Validator          *validator.Validate
	TransactionUseCase usecase.InterfaceTransactionUseCase
}

func NewTransactionController(log *logrus.Logger, val *validator.Validate, tu usecase.InterfaceTransactionUseCase) *TransactionController {
	return &TransactionController{
		Log:                log,
		Validator:          val,
		TransactionUseCase: tu,
	}
}

func (c *TransactionController) GetSpecificPickUp(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	transactionID := ctx.Params("transactionId")

	res, err := c.TransactionUseCase.SpecificPickUp(auth, transactionID)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Success{
		Message: "Success get Product Details",
		Data:    res,
		Status: response.Status{
			Code:    fiber.StatusOK,
			Message: "OK",
		},
	})
}

func (c *TransactionController) GetSpecificShop(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	transactionID := ctx.Params("transactionId")

	res, err := c.TransactionUseCase.SpecificShop(auth, transactionID)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Success{
		Message: "Success get Product Details",
		Data:    res,
		Status: response.Status{
			Code:    fiber.StatusOK,
			Message: "OK",
		},
	})
}

func (c *TransactionController) HistoryPickUp(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	res, err := c.TransactionUseCase.HistoryPickUp(auth)
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

func (c *TransactionController) HistoryShop(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	res, err := c.TransactionUseCase.HistoryShop(auth)
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
