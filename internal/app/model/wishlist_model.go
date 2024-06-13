package model

type MyWishlist struct {
	TotalCart int `json:"total_cart"`
	Product   []ExploreItems
}

type RequestManageWishlist struct {
	ProductID string `json:"product_id"`
}

type ResponseManageWishList struct {
	Message string `json:"message"`
}
