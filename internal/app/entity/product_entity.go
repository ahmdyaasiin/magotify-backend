package entity

type Product struct {
	ID                 string  `db:"id"`
	Name               string  `db:"name"`
	Description        string  `db:"description"`
	Quantity           int     `db:"quantity"`
	Price              float64 `db:"price"`
	DiscountPercentage int     `db:"discount_percentage"`
	Weight             float64 `db:"weight"`
	CreatedAt          int64   `db:"created_at"`
	UpdatedAt          int64   `db:"updated_at"`
	CategoryID         string  `db:"category_id"`
}

func (e Product) GetTableName() string {
	return "product"
}
