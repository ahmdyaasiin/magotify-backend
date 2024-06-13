package entity

type User struct {
	ID        string  `db:"id"`
	Name      string  `db:"name"`
	Email     string  `db:"email"`
	Password  string  `db:"password"`
	UrlPhoto  string  `db:"url_photo"`
	Balance   float64 `db:"balance"`
	CreatedAt int64   `db:"created_at"`
	UpdatedAt int64   `db:"updated_at"`
}

func (e User) GetTableName() string {
	return "users"
}
