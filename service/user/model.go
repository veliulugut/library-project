package user

import "time"

type CreateUserModel struct {
	Password        string `json:"password"`
	Email           string `json:"email"`
	UserName        string `json:"user_name"`
	ConfirmPassword string `json:"confirm_password"`
}

type GetUserModel struct {
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUserModel struct {
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	UpdatedAt time.Time `json:"updated_at" example:"2006-01-02T15:04:05Z"`
}
