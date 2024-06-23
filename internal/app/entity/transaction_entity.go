package entity

type Transaction struct {
	ID            string  `db:"id"`
	InvoiceNumber string  `db:"invoice_number"`
	TotalAmount   float64 `db:"total_amount"`
	ShippingCosts float64 `db:"shipping_costs"`
	Status        string  `db:"status"`
	ServiceName   string  `db:"service_name"`
	ServiceType   string  `db:"service_type"`
	ReceiptNumber string  `db:"receipt_number"`
	PaymentType   string  `db:"payment_type"`
	CreatedAt     int64   `db:"created_at"`
	UpdatedAt     int64   `db:"updated_at"`
	AddressID     string  `db:"address_id"`
	VoucherID     string  `db:"voucher_id"`
}

func (e Transaction) GetTableName() string {
	return "transactions"
}
