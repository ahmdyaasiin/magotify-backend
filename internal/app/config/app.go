package config

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http/middleware"
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
	addressRepository := repository.NewAddressRepository(config.DB)
	voucherRepository := repository.NewVoucherRepository(config.DB)
	cartRepository := repository.NewCartRepository(config.DB)
	bannerRepository := repository.NewBannerRepository(config.DB)
	wishlistRepository := repository.NewWishlistRepository(config.DB)
	productRepository := repository.NewProductRepository(config.DB)
	mediaRepository := repository.NewMediaRepository(config.DB)
	ratingRepository := repository.NewRatingRepository(config.DB)
	paymentRepository := repository.NewPaymentRepository(config.DB)
	categoryRepository := repository.NewCategoryRepository(config.DB)

	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, userRepository, addressRepository)
	menuUseCase := usecase.NewMenuUseCase(config.DB, config.Log, menuRepository, userRepository, addressRepository, voucherRepository, cartRepository, bannerRepository, categoryRepository)
	cartUseCase := usecase.NewCartUseCase(config.DB, config.Log, menuRepository, userRepository, cartRepository, productRepository)
	wishlistUseCase := usecase.NewWishlistUseCase(config.DB, config.Log, userRepository, wishlistRepository, cartRepository, productRepository)
	productUseCase := usecase.NewProductUseCase(config.DB, config.Log, productRepository, userRepository, cartRepository, mediaRepository, ratingRepository)
	paymentUseCase := usecase.NewPaymentUseCase(config.DB, config.Log, paymentRepository, userRepository, addressRepository)

	serverController := http.NewServerController()
	userController := http.NewUserController(config.Log, config.Validator, userUseCase)
	menuController := http.NewMenuController(config.Log, config.Validator, menuUseCase)
	cartController := http.NewCartController(config.Log, config.Validator, cartUseCase)
	wishlistController := http.NewWishlistController(config.Log, config.Validator, wishlistUseCase)
	productController := http.NewProductController(config.Log, config.Validator, productUseCase)
	paymentController := http.NewPaymentController(config.Log, config.Validator, paymentUseCase)

	userMiddleware := middleware.NewUserMiddleware(userUseCase)

	routeConfig := &route.Config{
		App:                config.App,
		ServerController:   serverController,
		UserController:     userController,
		MenuController:     menuController,
		CartController:     cartController,
		Middleware:         userMiddleware,
		WishlistController: wishlistController,
		ProductController:  productController,
		PaymentController:  paymentController,
	}

	routeConfig.Setup()
}
