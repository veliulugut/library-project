package book

import (
	"context"
	"fmt"
	"library/pkg/repository/dto"
	"library/pkg/repository/entadp"
)

var _ ServiceBook = (*Book)(nil)

func New(repo entadp.RepositoryInterface) *Book {
	return &Book{
		repo: repo,
	}
}

type Book struct {
	repo entadp.RepositoryInterface
}

func (b *Book) CreateBook(ctx context.Context, c *CreateBookModel) error {
	bookDTO := dto.Book{
		Title:     c.Title,
		Author:    c.Author,
		Genre:     c.Genre,
		Height:    c.Height,
		Publisher: c.Publisher,
	}

	if err := b.repo.Book().CreateBook(ctx, &bookDTO); err != nil {
		return fmt.Errorf("book srv / create book:%w", err)
	}

	return nil
}

func (b *Book) DeleteBook(ctx context.Context, id int) error {
	if err := b.repo.Book().DeleteBook(ctx, id); err != nil {
		return fmt.Errorf("book srv / delete book:%w", err)
	}

	return nil
}

func (b *Book) UpdatedBook(ctx context.Context, id int, c *UpdateBookModel) error {
	bookDTO := dto.Book{
		Title:     c.Title,
		Author:    c.Author,
		Genre:     c.Genre,
		Height:    c.Height,
		Publisher: c.Publisher,
	}

	if err := b.repo.Book().UpdatedBook(ctx, id, &bookDTO); err != nil {
		return fmt.Errorf("book srv / updated book :%w", err)
	}

	return nil
}

func (b *Book) GetBookByName(ctx context.Context, name string) (*dto.Book, error) {
	book, err := b.repo.Book().GetBookByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("book srv / get book by name :%w", err)
	}

	return book, nil
}

func (b *Book) GetBookByID(ctx context.Context, id int) (*dto.Book, error) {
	book, err := b.repo.Book().GetBookByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("book srv / get book by id :%w", err)
	}

	return book, nil
}

func (b *Book) ListBook(ctx context.Context, limit, offset int, orderBy string) ([]*dto.Book, error) {
	//TODO implement me
	panic("implement me")
}
