package book

import (
	"context"
	"fmt"
	"library/ent"
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

func (b *Book) GetBookByName(ctx context.Context, name string) (*GetBookByName, error) {
	var (
		dbBook *dto.Book
		err    error
	)

	if dbBook, err = b.repo.Book().GetBookByName(ctx, name); err != nil {
		return nil, fmt.Errorf("book srv / get book by name :%w", err)
	}

	return &GetBookByName{
		Title:     dbBook.Title,
		Author:    dbBook.Author,
		Genre:     dbBook.Genre,
		Height:    dbBook.Height,
		Publisher: dbBook.Publisher,
	}, nil
}

func (b *Book) GetBookByID(ctx context.Context, id int) (*GetBookModel, error) {
	var (
		dbBook *dto.Book
		err    error
	)

	if dbBook, err = b.repo.Book().GetBookByID(ctx, id); err != nil {
		return nil, fmt.Errorf("book srv / get book by id :%w", err)
	}

	return &GetBookModel{
		Title:     dbBook.Title,
		Author:    dbBook.Author,
		Genre:     dbBook.Genre,
		Height:    dbBook.Height,
		Publisher: dbBook.Publisher,
	}, nil
}

func (b *Book) ListBook(ctx context.Context, limit, offset int, orderBy string) ([]*GetBookModel, int, error) {
	books, count, err := b.repo.Book().ListBook(ctx, limit, offset, orderBy)
	if err != nil {
		return nil, 0, fmt.Errorf("book srv / list book: %w", err)
	}

	var bookModels []*GetBookModel
	for _, dbBook := range books {
		bookModel := dbBookToGetBookModel(dbBook)
		bookModels = append(bookModels, bookModel)
	}

	return bookModels, count, nil
}

func dbBookToGetBookModel(dbBook *ent.Book) *GetBookModel {
	return &GetBookModel{
		Title:     dbBook.Title,
		Author:    dbBook.Author,
		Genre:     dbBook.Genre,
		Height:    dbBook.Height,
		Publisher: dbBook.Publisher,
	}
}

func (b *Book) FetchBookData(limit, offset int, orderBy string, fetchAll bool) ([]*ent.Book, error) {
	books, err := b.repo.Book().FetchBookData(limit, offset, orderBy, fetchAll)
	if err != nil {
		return nil, fmt.Errorf("book srv / fetch book data :%w", err)
	}

	return books, nil
}
