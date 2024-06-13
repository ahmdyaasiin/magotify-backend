package model

import "github.com/ahmdyaasiin/magotify-backend/internal/app/entity"

type ExploreUser struct {
	ID           string  `json:"id"`
	FullName     string  `json:"full_name"`
	PhotoProfile string  `json:"photo_profile"`
	Balance      float64 `json:"balance"`
	Voucher      int     `json:"voucher"`
}

type ShopUser struct {
	ID    string `json:"id"`
	City  string `json:"city"`
	State string `json:"state"`
}

type ShopBanner struct {
	ID          string `json:"id"`
	UrlPhoto    string `json:"url_photo"`
	IsClickable string `json:"is_clickable"`
	Destination string `json:"destination"`
}

type Items struct {
	All   []ExploreItems
	Pupuk []ExploreItems
	Pakan []ExploreItems
}

type ExploreItems struct {
	ProductID     string  `json:"id" db:"product_id"`
	ProductImage  string  `json:"image" db:"product_image"`
	ProductName   string  `json:"name" db:"name"`
	RealPrice     string  `json:"real_price" db:"real_price"`
	DiscountPrice string  `json:"discount_price" db:"discount_price"`
	Rating        float64 `json:"rating" db:"average_rating"`
	Sold          int     `json:"sold" db:"sold"`
	IsWishlist    bool    `json:"is_wishlist" db:"is_wishlist"`
}

type ExploreTimeToCleanUp struct {
	VehicleID       string  `json:"vehicle_id" db:"id"`
	VehicleName     string  `json:"vehicle_name" db:"name"`
	UrlPhoto        string  `json:"url_photo" db:"url_photo"`
	Description     string  `json:"description" db:"description"`
	Status          string  `json:"status" db:"status"`
	WarehouseName   string  `json:"warehouse_name" db:"warehouse_name"`
	DistanceVehicle float64 `json:"distance_on_meter" db:"distance_m"`
}

type HotItemsSlice struct {
	Name     string         `json:"name" db:"name"`
	UrlPhoto string         `json:"url_photo" db:"url_photo"`
	Products []ExploreItems `json:"products"`
}

type ResponseMenuExplore struct {
	User          ExploreUser            `json:"user"`
	HotItems      []HotItemsSlice        `json:"hot_items"`
	TimeToCleanUp []ExploreTimeToCleanUp `json:"time_to_clean_up"`
}

type ResponseMenuShop struct {
	User            ShopUser        `json:"user"`
	TotalCart       int             `json:"total_cart"`
	Banner          []entity.Banner `json:"banners"`
	HotItems        []HotItemsSlice `json:"hot_items"`
	ProductDiscount []ExploreItems  `json:"products_discount"`
}
