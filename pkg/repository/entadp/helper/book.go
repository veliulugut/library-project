package helper

import (
	"library/ent"
	"library/pkg/repository/dto"
)

func DBBookToDTO(b *ent.Book) *dto.Book {
	return &dto.Book{
		ID:        (b.ID),
		Author:    b.Author,
		Genre:     b.Genre,
		Height:    b.Height,
		Title:     b.Title,
		Publisher: b.Publisher,
		UpdatedAt: b.UpdatedAt,
	}
}
