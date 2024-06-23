package model

type RequestUserRegister struct {
	FullName    string  `json:"full_name" db:"full_name" validate:"required"`
	Email       string  `json:"email" db:"email" validate:"required"`
	Password    string  `json:"password" db:"password" validate:"required"`
	PhoneNumber string  `json:"phone_number" db:"phone_number" validate:"required"`
	Address     string  `json:"address" db:"address" validate:"required"`
	District    string  `json:"district" db:"district" validate:"required"`
	City        string  `json:"city" db:"city" validate:"required"`
	State       string  `json:"state" db:"state" validate:"required"`
	PostalCode  string  `json:"postal_code" db:"postal_code" validate:"required"`
	Latitude    float64 `json:"latitude" db:"latitude" validate:"required"`
	Longitude   float64 `json:"longitude" db:"longitude" validate:"required"`
}

type RequestUserLogin struct {
	Email    string `json:"email" db:"email" validate:"required"`
	Password string `json:"password" db:"password" validate:"required"`
}

type RequestUserForgotPassword struct {
	Email string `json:"email" db:"email"`
}

type UUIDMiddleware struct {
	UUID string `json:"uuid"`
}

type ResponseUser struct {
	Token string `json:"token"`
}
