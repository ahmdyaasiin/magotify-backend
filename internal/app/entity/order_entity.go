package entity

type Order struct {
	ID            string  `json:"id" db:"id"`
	InvoiceNumber string  `json:"invoice_number" db:"invoice_number"`
	TotalAmount   float64 `json:"total_amount" db:"total_amount"`
	Weight        float64 `json:"weight" db:"weight"`
	Status        string  `json:"status" db:"status"`
	PaymentType   string  `json:"payment_type" db:"payment_type"`
	CreatedAt     int64   `json:"created_at" db:"created_at"`
	UpdatedAt     int64   `json:"updated_at" db:"updated_at"`
	AddressID     string  `json:"address_id" db:"address_id"`
	DriverID      string  `json:"driver_id" db:"driver_id"`
	VoucherID     string  `json:"voucher_id" db:"voucher_id"`
}

func (e Order) GetTableName() string {
	return "orders"
}
