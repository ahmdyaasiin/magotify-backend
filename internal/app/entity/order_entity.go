package entity

type Order struct {
	ID            string  `json:"id"`
	InvoiceNumber string  `json:"invoice_number"`
	TotalAmount   float64 `json:"total_amount"`
	Weight        float64 `json:"weight"`
	Status        string  `json:"status"`
	PaymentType   string  `json:"payment_type"`
	CreatedAt     int64   `json:"created_at"`
	UpdatedAt     int64   `json:"updated_at"`
	AddressID     string  `json:"address_id"`
	DriverID      string  `json:"driver_id"`
	VoucherID     string  `json:"voucher_id"`
}

func (e Order) GetTableName() string {
	return "orders"
}
