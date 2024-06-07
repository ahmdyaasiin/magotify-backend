package entity

type Media struct {
	ID        string `db:"id"`
	UrlPhoto  string `db:"url_photo"`
	CreatedAt int64  `db:"created_at"`
	ProductID string `db:"product_id"`
}

func (e Media) GetTableName() string {
	return "media"
}
