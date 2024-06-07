package entity

type Vehicle struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	UrlPhoto    string `db:"url_photo"`
	Status      string `db:"status"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	WarehouseID string `db:"warehouse_id"`
}

func (e Vehicle) GetTableName() string {
	return "vehicles"
}
