//go:generate mockgen -destination=service_mocks.go -package=book library/service/book ServiceBook
package book

import (
	"context"
	"library/ent"
)

type ServiceBook interface {
	CreateBook(ctx context.Context, c *CreateBookModel) error
	DeleteBook(ctx context.Context, id int) error
	UpdatedBook(ctx context.Context, id int, c *UpdateBookModel) error
	GetBookByName(ctx context.Context, name string) (*GetBookByName, error)
	ListBook(ctx context.Context, limit, offset int, orderBy string) ([]*GetBookModel, int, error)
	GetBookByID(ctx context.Context, id int) (*GetBookModel, error)
	FetchBookData(limit, offset int, orderBy string, fetchAll bool) ([]*ent.Book, error)
}
