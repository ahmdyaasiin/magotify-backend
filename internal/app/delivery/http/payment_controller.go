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

func (c *PaymentController) ValidatePaymentPickUp(ctx *fiber.Ctx) error {

	request := new(model.RequestValidatePayment)

	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.Validator.Struct(request); err != nil {
		return err
	}

	res, err := c.PaymentUseCase.ValidatePickUp(request)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Success{
		Message: "Success validate Payment",
		Data:    res,
		Status: response.Status{
			Code:    fiber.StatusOK,
			Message: "OK",
		},
	})
}

func (c *PaymentController) ValidatePaymentShop(ctx *fiber.Ctx) error {
	fmt.Println("hadhe")
	request := new(model.RequestValidatePayment)

	fmt.Println("hadhe2s")
	fmt.Println(request)

	if err := ctx.BodyParser(request); err != nil {
		fmt.Println("error nih cok" + err.Error())
		return fiber.ErrBadRequest
	}

	if err := c.Validator.Struct(request); err != nil {
		fmt.Println("error nih wokwok" + err.Error())
		return err
	}

	fmt.Println("passed")

	res, err := c.PaymentUseCase.ValidateShop(request)
	if err != nil {
		fmt.Println("error hadeh: " + err.Error())
		return err
	}

	return ctx.JSON(response.Success{
		Message: "Success validate Payment",
		Data:    res,
		Status: response.Status{
			Code:    fiber.StatusOK,
			Message: "OK",
		},
	})
}

func (c *PaymentController) CreatePaymentPickUp(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.RequestCreatePickUp)

	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.Validator.Struct(request); err != nil {
		return err
	}

	res, err := c.PaymentUseCase.CreatePickUp(auth, request)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Success{
		Message: "Success Create Pick Up",
		Data:    res,
		Status: response.Status{
			Code:    fiber.StatusOK,
			Message: "OK",
		},
	})
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

func (c *PaymentController) GetPaymentPickUp(ctx *fiber.Ctx) error {

	// address_id (optional)
	// warehouse_id
	// vehicles_id (optional)
	// jenis_sampah (optional?)
	// berat_sampah (optional?)
	// voucher_id (optional)

	// recent address/primary address
	// addresses except recent address
	// warehouse details
	// jenis jenis sampah
	// berat sampah
	// jenis jenis kendaraan
	// recent voucher
	// vouchers

	auth := middleware.GetUser(ctx)

	addressID := ctx.Query("address_id")
	warehouseID := ctx.Query("warehouse_id")
	vehiclesID := ctx.Query("vehicles_id")
	//wasteType := ctx.Query("waste_type")
	//wasteWeight := ctx.Query("waste_weight")
	voucherID := ctx.Query("voucher_id")

	if warehouseID == "" {
		return ctx.JSON(response.Message{Message: "warehouse_id mana?"})
	}

	res, err := c.PaymentUseCase.PickUp(auth, warehouseID, addressID, vehiclesID, voucherID)
	if err != nil {
		return err
	}

	return ctx.JSON(response.Success{
		Message: "Success get Payment Pick Up",
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

	voucherID := ctx.Query("voucher_id")

	// gapake voucher if empty

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

	res, err := c.PaymentUseCase.Shop(auth, productIds, addressID, quantites, voucherID)
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
