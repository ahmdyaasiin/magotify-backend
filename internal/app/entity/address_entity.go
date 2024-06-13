package entity

type Address struct {
	ID          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Address     string  `json:"address" db:"address"`
	District    string  `json:"district" db:"district"`
	City        string  `json:"city" db:"city"`
	State       string  `json:"state" db:"state"`
	PostalCode  string  `json:"postal_code" db:"postal_code"`
	Latitude    float64 `json:"latitude" db:"latitude"`
	Longitude   float64 `json:"longitude" db:"longitude"`
	IsPrimary   bool    `json:"is_primary" db:"is_primary"`
	PhoneNumber string  `json:"phone_number" db:"phone_number"`
	CreatedAt   int64   `json:"-" db:"created_at"`
	UpdatedAt   int64   `json:"-" db:"updated_at"`
	UserID      string  `json:"-" db:"user_id"`
}

func (e Address) GetTableName() string {
	return "addresses"
}
