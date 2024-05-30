package http

import "github.com/gofiber/fiber/v2"

type ServerController struct {
	//
}

func NewServerController() *ServerController {
	return &ServerController{}
}

func (c *ServerController) Status(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status": "OK",
	})
}
