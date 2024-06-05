package entity

type User struct {
	ID           string `db:"id"`
	Name         string `db:"name"`
	Email        string `db:"email"`
	Password     string `db:"password"`
	PhoneNumber  string `db:"phone_number"`
	PhotoProfile string `db:"photo_profile"`
	CreatedAt    int64  `db:"created_at"`
	UpdatedAt    int64  `db:"updated_at"`
}

func (u User) GetTableName() string {
	return "users"
}
