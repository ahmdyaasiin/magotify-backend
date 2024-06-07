package entity

type Warehouse struct {
	ID        string  `db:"id"`
	Name      string  `db:"name"`
	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
	CreatedAt int64   `db:"created_at"`
	UpdatedAt int64   `db:"updated_at"`
}

func (e Warehouse) GetTableName() string {
	return "warehouses"
}
