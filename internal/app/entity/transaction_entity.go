package entity

type Transaction struct {
	ID          string  `db:"id"`
	TotalAmount float64 `db:"total_amount"`
	Status      string  `db:"status"`
	CreatedAt   int64   `db:"created_at"`
	UpdatedAt   int64   `db:"updated_at"`
	UserID      string  `db:"user_id"`
}

func (e Transaction) GetTableName() string {
	return "transactions"
}
