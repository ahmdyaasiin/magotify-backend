package entity

type TransactionItem struct {
	ID            string  `db:"id"`
	Quantity      int     `db:"quantity"`
	TotalPrice    float64 `db:"total_price"`
	TransactionID string  `db:"transaction_id"`
	ProductID     string  `db:"product_id"`
	CreatedAt     int64   `db:"created_at"`
}

func (e TransactionItem) GetTableName() string {
	return "transaction_items"
}
