package middleware

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/usecase"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/firebase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strings"
)

func NewUserMiddleware(u usecase.InterfaceUserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if ctx.Path() == "/v1/payment/shop/validate" {
			return ctx.Next()
		}

		bearerArray := ctx.GetReqHeaders()["Authorization"]

		var bearer string
		if len(bearerArray) > 0 {
			bearer = bearerArray[0]
		} else {
			return fiber.ErrBadRequest
		}

		if strings.HasPrefix(bearer, "Bearer ") {
			bearer = strings.Split(bearer, " ")[1]
		} else {
			return fiber.ErrBadRequest
		}

		uid, err := firebase.DecodeCustomToken(bearer)
		if err != nil {
			fmt.Println("error nih: " + err.Error())
			return err
		}

		err = u.Verify(&model.UUIDMiddleware{UUID: uid})
		if err != nil {
			return err
		}

		ctx.Locals("user_id", uid)
		return ctx.Next()
	}
}

func NewCorsMiddleware() fiber.Handler {
	return cors.New()
}

func GetUser(ctx *fiber.Ctx) string {
	return ctx.Locals("user_id").(string)
}
