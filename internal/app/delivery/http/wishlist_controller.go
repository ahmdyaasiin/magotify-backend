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

type WishlistController struct {
	Log             *logrus.Logger
	Validator       *validator.Validate
	WishlistUseCase usecase.InterfaceWishlistUseCase
}

func NewWishlistController(log *logrus.Logger, val *validator.Validate, wu usecase.InterfaceWishlistUseCase) *WishlistController {
	return &WishlistController{
		Log:             log,
		Validator:       val,
		WishlistUseCase: wu,
	}
}

func (c *WishlistController) GetWishlist(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	res, err := c.WishlistUseCase.GetWishlist(auth)
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

func (c *WishlistController) ManageWishlist(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.RequestManageWishlist)

	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	if err := c.Validator.Struct(request); err != nil {
		return err
	}

	// usecase
	res, err := c.WishlistUseCase.AddWishlist(auth, request)
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
