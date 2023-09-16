package dto

import "time"

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Genre     string    `json:"genre"`
	Height    string    `json:"height"`
	Publisher string    `json:"publisher"`
	CreatedAt time.Time `json:"created_at" example:"2006-01-02T15:04:05Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2006-01-02T15:04:05Z"`
}
