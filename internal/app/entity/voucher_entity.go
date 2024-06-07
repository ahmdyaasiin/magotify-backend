package entity

type Voucher struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Status      bool   `db:"status"`
	Amount      int    `db:"amount"`
	IsPercent   bool   `db:"is_percent"`
	UrlLogo     string `db:"url_logo"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	UserID      string `db:"user_id"`
}

func (e Voucher) GetTableName() string {
	return "vouchers"
}
