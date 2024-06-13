package route

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	//
	App                *fiber.App
	Middleware         fiber.Handler
	ServerController   *http.ServerController
	UserController     *http.UserController
	MenuController     *http.MenuController
	CartController     *http.CartController
	WishlistController *http.WishlistController
	ProductController  *http.ProductController
	PaymentController  *http.PaymentController
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
	//menu.Get("pick-up")
	menu.Get("shop", c.MenuController.Shop)
	//menu.Get("transaction")

	product := v1.Group("/product")
	product.Get(":productId/details", c.ProductController.GetProductDetails)

	payment := v1.Group("/payment")
	//payment.Get("pick-up")
	payment.Post("validate", c.PaymentController.ValidatePayment)
	payment.Get("shop", c.PaymentController.GetPaymentShop)
	payment.Post("shop/create", c.PaymentController.CreatePaymentShop)

	//transaction := v1.Group("/transaction")
	//transaction.Get("pick-up")
	//transaction.Get("shop")
	//transaction.Get("details")

	user := v1.Group("/user")
	user.Get("cart", c.CartController.GetCart)
	user.Post("cart/manage", c.CartController.AddToCart)
	user.Get("wishlist", c.WishlistController.GetWishlist)
	user.Post("wishlist/manage", c.WishlistController.ManageWishlist)

}
