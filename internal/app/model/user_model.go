package model

type RequestUserRegister struct {
	FullName    string `json:"full_name" db:"full_name" validate:"required"`
	Email       string `json:"email" db:"email" validate:"required"`
	Password    string `json:"password" db:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" db:"phone_number" validate:"required"`
}

type RequestUserLogin struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type RequestUserForgotPassword struct {
	Email string `json:"email" db:"email"`
}

type ResponseUser struct {
	Token string `json:"token"`
}
