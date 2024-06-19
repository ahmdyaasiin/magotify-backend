package model

type MyCart struct {
	Product   []ProductCart   `json:"product"`
	TotalCart int             `json:"total_cart"`
	HotItems  []HotItemsSlice `json:"hot_items"`
}

type ProductCart struct {
	ID            string  `json:"id" db:"id"`
	UrlPhoto      string  `json:"url_photo" db:"url_photo"`
	Name          string  `json:"name" db:"name"`
	CategoryName  string  `json:"category_name" db:"cat_name"`
	QuantityCart  int     `json:"quantity_cart" db:"quantity"`
	Weight        float64 `json:"weight" db:"weight"`
	RealPrice     float64 `json:"real_price" db:"price"`
	DiscountPrice float64 `json:"discount_price" db:"discount_price"`
}

type RequestAddCart struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity" validate:"gte=0"`
}

type ResponseAddCart struct {
	TotalCart int `json:"total_cart"`
}
