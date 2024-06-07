package entity

type Cart struct {
	ID        string `db:"id"`
	Quantity  int    `db:"quantity"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
	UserID    string `db:"user_id"`
	ProductID string `db:"product_id"`
}

func (e Cart) GetTableName() string {
	return "carts"
}
