package model

type ExploreUser struct {
	FullName     string `json:"full_name"`
	PhotoProfile string `json:"photo_profile"`
	Wallet       string `json:"wallet"`
	Voucher      int    `json:"voucher"`
}

type ExploreItems struct {
	PhotoProduct  string     `json:"photo_product"`
	ProductName   string     `json:"product_name"`
	RealPrice     string     `json:"real_price"`
	DiscountPrice string     `json:"discount_price"`
	Rating        ItemRating `json:"rating"`
}

type ItemRating struct {
	Star int `json:"star"`
	Sold int `json:"sold"`
}

type ExploreTimeToCleanUp struct {
	PhotoTTC       string `json:"photo_ttc"`
	NameTTC        string `json:"name_ttc"`
	DescriptionTTC string `json:"description_ttc"`
	StatusTTC      string `json:"status_ttc"`
	DistanceTTC    string `json:"distance_ttc"`
}

type ResponseMenuExplore struct {
	User          ExploreUser
	HotItems      ExploreItems
	TimeToCleanUp ExploreTimeToCleanUp
}
