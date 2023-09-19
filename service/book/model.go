package book

import "library/ent"

type CreateBookModel struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	Height    string `json:"height"`
	Publisher string `json:"publisher"`
}

type UpdateBookModel struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	Height    string `json:"height"`
	Publisher string `json:"publisher"`
}

type GetBookModel struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	Height    string `json:"height"`
	Publisher string `json:"publisher"`
}

type GetBookByName struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	Height    string `json:"height"`
	Publisher string `json:"publisher"`
}

type ListBookResponse struct {
	RowCount int
	Data     []*ent.Book
}
