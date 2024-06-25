package entity

type Banner struct {
	ID          string `json:"id"`
	UrlPhoto    string `json:"url_photo" db:"url_photo"`
	IsClickable bool   `json:"is_clickable" db:"is_clickable"`
	Destination string `json:"destination"`
	CreatedAt   int64  `json:"-" db:"created_at"`
	UpdatedAt   int64  `json:"-" db:"updated_at"`
}
