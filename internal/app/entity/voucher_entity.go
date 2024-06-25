package entity

type Voucher struct {
	ID          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Status      bool    `json:"status" db:"status" rules:"can_zero"`
	Amount      int     `json:"amount" db:"amount"`
	IsPercent   bool    `json:"is_percent" db:"is_percent"`
	UrlLogo     string  `json:"url_logo" db:"url_logo"`
	MinAmount   float64 `json:"min_amount" db:"min_amount"`
	CreatedAt   int64   `json:"-" db:"created_at"`
	UpdatedAt   int64   `json:"-" db:"updated_at"`
	UserID      string  `json:"-" db:"user_id"`
}

func (e Voucher) GetTableName() string {
	return "vouchers"
}
