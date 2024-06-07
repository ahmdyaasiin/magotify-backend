package entity

type Wishlist struct {
	ID        string `db:"id"`
	CreatedAt int64  `db:"created_at"`
	UserID    string `db:"user_id"`
	ProductID string `db:"product_id"`
}

func (e Wishlist) GetTableName() string {
	return "wishlists"
}
