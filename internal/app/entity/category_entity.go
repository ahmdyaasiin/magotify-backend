package entity

type Category struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	UrlPhoto  string `db:"url_photo"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

func (e Category) GetTableName() string {
	return "categories"
}
