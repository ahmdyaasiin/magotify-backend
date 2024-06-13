package model

type MediaProduct struct {
	ID       string `json:"id" db:"id"`
	UrlPhoto string `json:"url_photo" db:"url_photo"`
}

type ReviewProduct struct {
	Name      string `json:"name" db:"name"`
	UrlPhoto  string `json:"url_photo" db:"url_photo"`
	Star      int    `json:"star" db:"star"`
	Content   string `json:"content" db:"content"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
}

type PD struct {
	ID             string  `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	RealPrice      float64 `json:"real_price" db:"price"`
	DiscountPrice  float64 `json:"discount_price" db:"discount_price"`
	Star           float64 `json:"star" db:"average_rating"`
	TotalSold      int     `json:"total_sold" db:"total_sold"`
	RecentQuantity int     `json:"recent_quantity" db:"cart_quantity"`
	Description    string  `json:"description" db:"description"`
	IsWishlist     bool    `json:"is_wishlist" db:"is_wishlist"`

	// media
	Media []MediaProduct `json:"media"`

	// review
	Review string `json:"review"` // please fix this

	// discus
	Discuss string `json:"discuss"` // please fix this
}

type ProductDetails struct {
	PD              PD             `json:"product_details"`
	TotalCart       int            `json:"total_cart"`
	ProductDiscount []ExploreItems `json:"products_discount"`
}
