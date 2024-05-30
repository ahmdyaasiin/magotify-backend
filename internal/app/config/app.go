package config

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http/route"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	App       *fiber.App
	DB        *sqlx.DB
	Log       *logrus.Logger
	Validator *validator.Validate
}

func App(config *AppConfig) {
	//

	serverController := http.NewServerController()

	routeConfig := &route.Config{
		App:              config.App,
		ServerController: serverController,
	}

	routeConfig.Setup()
}
