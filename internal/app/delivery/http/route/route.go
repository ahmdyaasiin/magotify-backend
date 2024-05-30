package route

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	//
	App              *fiber.App
	ServerController *http.ServerController
}

func (c *Config) Setup() {
	c.V1()
}

func (c *Config) V1() {
	// grouping
	v1 := c.App.Group("/v1")
	v1.Get("status", c.ServerController.Status)

}
