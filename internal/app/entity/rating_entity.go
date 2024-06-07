package entity

type Rating struct {
	ID            string `db:"id"`
	Content       string `db:"content"`
	Star          int    `db:"star"`
	CreatedAt     int64  `db:"created_at"`
	UpdatedAt     int64  `db:"updated_at"`
	UserID        string `db:"user_id"`
	TransactionID string `db:"transaction_id"`
}

func (e Rating) GetTableName() string {
	return "ratings"
}
