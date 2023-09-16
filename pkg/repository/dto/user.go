package dto

import "time"

type User struct {
	ID        uint8     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at" example:"2006-01-02T15:04:05Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2006-01-02T15:04:05Z"`
}

type UpdateUserModel struct {
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	UpdatedAt time.Time `json:"updated_at" example:"2006-01-02T15:04:05Z"`
}
