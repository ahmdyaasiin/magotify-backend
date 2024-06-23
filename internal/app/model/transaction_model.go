package model

type ResponseTransactionPickUp struct {
	//
	DriverName  string `json:"driver_name"`
	Address     string `json:"address"`
	TotalAmount string `json:"total_amount"`
	Weight      string `json:"weight"`
	Distance    string `json:"distance"`
	CreatedAt   string `json:"created_at"`
}

type ResponseTransactionShop struct {
	//
	InvoiceNumber      string `json:"invoice_number"`
	ProductImage       string `json:"product_image"`
	ProductName        string `json:"product_name"`
	ProductQuantity    string `json:"product_quantity"`
	TotalProduct       string `json:"total_product"`
	AddressName        string `json:"address_name"`
	TotalWeightProduct string `json:"total_weight_product"`
	CreatedAt          string `json:"created_at"`
	TotalPrice         string `json:"total_price"`
}

type ResponseSpecificTransactionShop struct {
	//

}

type ResponseSpecificTransactionPickUp struct {
	//

}
