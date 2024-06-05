package config

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http/route"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/repository"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/usecase"
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

	userRepository := repository.NewUserRepository(config.DB)
	menuRepository := repository.NewMenuRepository(config.DB)

	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, userRepository)
	menuUseCase := usecase.NewMenuUseCase(config.DB, config.Log, menuRepository)

	serverController := http.NewServerController()
	userController := http.NewUserController(config.Log, config.Validator, userUseCase)
	menuController := http.NewMenuController(config.Log, config.Validator, menuUseCase)

	routeConfig := &route.Config{
		App:              config.App,
		ServerController: serverController,
		UserController:   userController,
		MenuController:   menuController,
	}

	routeConfig.Setup()
}
