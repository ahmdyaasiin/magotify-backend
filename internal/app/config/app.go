package config

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http/middleware"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http/route"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/repository"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/usecase"
	"github.com/ahmdyaasiin/magotify-backend/internal/scheduler"
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
	warehouseRepository := repository.NewWarehouseRepository(config.DB)
	vehiclesRepository := repository.NewVehicleRepository(config.DB)
	driverRepository := repository.NewDriverRepository(config.DB)
	transactionRepository := repository.NewTransactionRepository(config.DB)
	transactionItemRepository := repository.NewTransactionItemRepository(config.DB)
	orderRepository := repository.NewOrderRepository(config.DB)

	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, userRepository, addressRepository)
	menuUseCase := usecase.NewMenuUseCase(config.DB, config.Log, menuRepository, userRepository, addressRepository, voucherRepository, cartRepository, bannerRepository, categoryRepository)
	cartUseCase := usecase.NewCartUseCase(config.DB, config.Log, menuRepository, userRepository, cartRepository, productRepository, categoryRepository)
	wishlistUseCase := usecase.NewWishlistUseCase(config.DB, config.Log, userRepository, wishlistRepository, cartRepository, productRepository)
	productUseCase := usecase.NewProductUseCase(config.DB, config.Log, productRepository, userRepository, cartRepository, mediaRepository, ratingRepository)
	paymentUseCase := usecase.NewPaymentUseCase(config.DB, config.Log, paymentRepository, userRepository, addressRepository, voucherRepository, warehouseRepository, vehiclesRepository, driverRepository, productRepository, transactionRepository, transactionItemRepository, orderRepository)
	transactionUseCase := usecase.NewTransactionUseCase(config.DB, config.Log, transactionRepository, userRepository, orderRepository, paymentRepository)

	serverController := http.NewServerController()
	userController := http.NewUserController(config.Log, config.Validator, userUseCase)
	menuController := http.NewMenuController(config.Log, config.Validator, menuUseCase)
	cartController := http.NewCartController(config.Log, config.Validator, cartUseCase)
	wishlistController := http.NewWishlistController(config.Log, config.Validator, wishlistUseCase)
	productController := http.NewProductController(config.Log, config.Validator, productUseCase)
	paymentController := http.NewPaymentController(config.Log, config.Validator, paymentUseCase)
	transactionController := http.NewTransactionController(config.Log, config.Validator, transactionUseCase)

	userMiddleware := middleware.NewUserMiddleware(userUseCase)
	corsMiddleware := middleware.NewCorsMiddleware()

	routeConfig := &route.Config{
		App:                   config.App,
		ServerController:      serverController,
		UserController:        userController,
		MenuController:        menuController,
		CartController:        cartController,
		Middleware:            userMiddleware,
		Cors:                  corsMiddleware,
		WishlistController:    wishlistController,
		ProductController:     productController,
		PaymentController:     paymentController,
		TransactionController: transactionController,
	}

	scheduler.Run(transactionUseCase)
	routeConfig.Setup()
}
