package http

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http/middleware"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/usecase"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strings"
)

type PaymentController struct {
	Log            *logrus.Logger
	Validator      *validator.Validate
	PaymentUseCase usecase.InterfacePaymentUseCase
}

func NewPaymentController(log *logrus.Logger, val *validator.Validate, pu usecase.InterfacePaymentUseCase) *PaymentController {
	return &PaymentController{
		Log:            log,
		Validator:      val,
		PaymentUseCase: pu,
	}
}

func (c *PaymentController) ValidatePayment(ctx *fiber.Ctx) error {

	return nil
}

func (c *PaymentController) CreatePaymentShop(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.RequestCreatePayment)

	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.Validator.Struct(request); err != nil {
		return err
	}

	fmt.Println(request)

	// usecase
	res, err := c.PaymentUseCase.CreateShop(auth, request)
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

func (c *PaymentController) GetPaymentShop(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	addressID := ctx.Query("address_id")

	// pake main address kalo gaada addressID nya
	//if addressID == "" {
	//	return ctx.JSON(response.Message{Message: "Address ID nya mana"})
	//}

	productIdsParam := ctx.Query("product_ids")
	productIds := strings.Split(productIdsParam, ",")

	// validasi uuid
	if len(productIds) == 1 && productIds[0] == "" {
		return ctx.JSON(response.Message{Message: "Product IDs nya mana"})
	}

	quantitesParam := ctx.Query("quantites")
	if strings.Contains(quantitesParam, "-") {
		return ctx.JSON(response.Message{Message: "Quantity tidak boleh negatif"})
	}

	quantites := strings.Split(quantitesParam, ",")

	if len(productIds) != len(quantites) {
		return ctx.JSON(response.Message{Message: "Product dan Quantitynya tidak sama"})
	}

	// add validasi jika quantity mines

	res, err := c.PaymentUseCase.Shop(auth, productIds, addressID, quantites)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Success{
		Message: "Success get Payment Shop",
		Data:    res,
		Status: response.Status{
			Code:    fiber.StatusOK,
			Message: "OK",
		},
	})
}
