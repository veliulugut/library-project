package book

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
