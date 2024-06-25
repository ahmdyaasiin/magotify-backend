package model

type ResponseTransactionPickUp struct {
	//
	ID          string `json:"id" db:"id"`
	DriverName  string `json:"driver_name" db:"name"`
	Address     string `json:"address" db:"address"`
	TotalAmount string `json:"total_amount" db:"total_amount"`
	Weight      string `json:"weight" db:"weight"`
	Distance    string `json:"distance" db:"distance_m"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

type ResponseTransactionShop struct {
	//
	ID                 string `json:"id" db:"id"`
	InvoiceNumber      string `json:"invoice_number" db:"invoice_number"`
	ProductImage       string `json:"product_image" db:"product_image"`
	ProductName        string `json:"product_name" db:"product_name"`
	ProductQuantity    string `json:"product_quantity" db:"product_quantity"`
	TotalProduct       string `json:"total_product" db:"total_products"`
	AddressName        string `json:"address_name" db:"address_name"`
	TotalWeightProduct string `json:"total_weight_product" db:"total_weight"`
	CreatedAt          string `json:"created_at" db:"created_at"`
	TotalPrice         string `json:"total_price" db:"total_price"`
}

type ResponseSpecificTransactionShop struct {
	//
	Products         []ProductCart `json:"products"`
	TransactionID    string        `json:"transaction_id" db:"transaction_id"`
	InvoiceNumber    string        `json:"invoice_number" db:"invoice_number"`
	TotalAmount      string        `json:"total_amount" db:"total_amount"`
	ShippingCosts    string        `json:"shipping_costs" db:"shipping_costs"`
	Status           string        `json:"status"`
	ServiceType      string        `json:"service_type" db:"service_type"`
	ServiceName      string        `json:"service_name" db:"service_name"`
	ReceiptNumber    string        `json:"receipt_number" db:"receipt_number"`
	PaymentType      string        `json:"payment_type" db:"payment_type"`
	CreatedAt        string        `json:"created_at" db:"created_at"`
	Address          string        `json:"address"`
	AddressName      string        `json:"address_name" db:"address_name"`
	VoucherID        string        `json:"voucher_id" db:"voucher_id"`
	VoucherName      string        `json:"voucher_name" db:"voucher_name"`
	VoucherAmount    string        `json:"voucher_amount" db:"voucher_amount"`
	VoucherIsPercent string        `json:"voucher_is_percent" db:"voucher_is_percent"`
	VoucherUrlLogo   string        `json:"voucher_url_logo" db:"voucher_url_logo"`
}

type ResponseSpecificTransactionPickUp struct {
	//
	TransactionID      string `json:"transaction_id" db:"transaction_id"`
	InvoiceNumber      string `json:"invoice_number" db:"invoice_number"`
	TotalAmount        string `json:"total_amount" db:"total_amount"`
	Weight             string `json:"weight" db:"weight"`
	PlateNumber        string `json:"plate_number" db:"plate_number"`
	VehicleName        string `json:"vehicle_name" db:"vehicle_name"`
	VehicleStatus      string `json:"vehicle_status" db:"vehicle_status"`
	WarehouseName      string `json:"warehouse_name" db:"warehouse_name"`
	WarehouseAddress   string `json:"warehouse_address" db:"warehouse_address"`
	VehicleDuration    string `json:"vehicle_duration" db:"vehicle_duration"`
	VehicleDescription string `json:"vehicle_description" db:"vehicle_description"`
	AddressName        string `json:"address_name" db:"address_name"`
	UserAddress        string `json:"user_address" db:"user_address"`
	DistanceM          string `json:"distance_m" db:"distance_m"`
}
