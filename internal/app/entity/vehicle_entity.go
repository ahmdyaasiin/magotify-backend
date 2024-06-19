package entity

type Vehicle struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Duration    string `json:"duration" db:"duration"`
	Description string `json:"description" db:"description"`
	UrlPhoto    string `json:"url_photo" db:"url_photo"`
	Status      string `json:"status" db:"status"`
	CreatedAt   int64  `json:"-" db:"created_at"`
	UpdatedAt   int64  `json:"-" db:"updated_at"`
}

func (e Vehicle) GetTableName() string {
	return "vehicles"
}
