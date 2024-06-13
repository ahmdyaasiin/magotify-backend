package http

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http/middleware"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/usecase"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductController struct {
	Log            *logrus.Logger
	Validator      *validator.Validate
	ProductUseCase usecase.InterfaceProductUseCase
}

func NewProductController(log *logrus.Logger, val *validator.Validate, pu usecase.InterfaceProductUseCase) *ProductController {
	return &ProductController{
		Log:            log,
		Validator:      val,
		ProductUseCase: pu,
	}
}

func (c *ProductController) GetProductDetails(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	productID := ctx.Params("productId")

	res, err := c.ProductUseCase.GetDetails(auth, productID)
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
