package route

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	//
	App                   *fiber.App
	Middleware            fiber.Handler
	ServerController      *http.ServerController
	UserController        *http.UserController
	MenuController        *http.MenuController
	CartController        *http.CartController
	WishlistController    *http.WishlistController
	ProductController     *http.ProductController
	PaymentController     *http.PaymentController
	TransactionController *http.TransactionController
}

func (c *Config) Setup() {
	c.V1()
}

func (c *Config) V1() {
	// grouping
	v1 := c.App.Group("/v1")
	v1.Get("status", c.ServerController.Status)

	auth := v1.Group("/auth")
	auth.Post("register", c.UserController.Register)
	auth.Post("login", c.UserController.Login)

	v1.Use(c.Middleware)

	menu := v1.Group("/menu")
	menu.Get("explore", c.MenuController.Explore)
	menu.Get("shop", c.MenuController.Shop)

	product := v1.Group("/product")
	product.Get(":productId/details", c.ProductController.GetProductDetails)

	payment := v1.Group("/payment")
	payment.Get("shop", c.PaymentController.GetPaymentShop)
	payment.Get("pick_up", c.PaymentController.GetPaymentPickUp)
	payment.Post("shop/create", c.PaymentController.CreatePaymentShop)
	payment.Post("pick_up/create", c.PaymentController.CreatePaymentPickUp)
	payment.Post("shop/validate", c.PaymentController.ValidatePaymentShop)
	payment.Post("pick_up/validate", c.PaymentController.ValidatePaymentPickUp)

	transaction := v1.Group("/transaction")
	transaction.Get("pick_up", c.TransactionController.HistoryPickUp)
	transaction.Get("shop", c.TransactionController.HistoryShop)
	transaction.Get("pick_up/:transactionId", c.TransactionController.GetSpecificPickUp)
	transaction.Get("shop/:transactionId", c.TransactionController.GetSpecificShop)

	user := v1.Group("/user")
	user.Get("cart", c.CartController.GetCart)
	user.Post("cart/manage", c.CartController.AddToCart)
	user.Get("wishlist", c.WishlistController.GetWishlist)
	user.Post("wishlist/manage", c.WishlistController.ManageWishlist)

}
