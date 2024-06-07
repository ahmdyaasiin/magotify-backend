package entity

type Address struct {
	ID         string  `json:"id" db:"id"`
	Name       string  `json:"name" db:"name"`
	Address    string  `json:"address" db:"address"`
	City       string  `json:"city" db:"city"`
	State      string  `json:"state" db:"state"`
	PostalCode string  `json:"postal_code" db:"postal_code"`
	Latitude   float64 `json:"latitude" db:"latitude"`
	Longitude  float64 `json:"longitude" db:"longitude"`
	IsPrimary  bool    `json:"is_primary" db:"is_primary"`
	CreatedAt  int64   `json:"created_at" db:"created_at"`
	UpdatedAt  int64   `json:"updated_at" db:"updated_at"`
	UserID     string  `json:"user_id" db:"user_id"`
}

func (e Address) GetTableName() string {
	return "addresses"
}
