//go:generate mockgen -destination=service_mocks.go -package=book library/service/book ServiceBook
package book

import (
	"context"
	"library/pkg/repository/dto"
)

type ServiceBook interface {
	CreateBook(ctx context.Context, c *CreateBookModel) error
	DeleteBook(ctx context.Context, id int) error
	UpdatedBook(ctx context.Context, id int, c *UpdateBookModel) error
	GetBookByName(ctx context.Context, name string) (*dto.Book, error)
	ListBook(ctx context.Context, limit, offset int, orderBy string) ([]*dto.Book, error)
	GetBookByID(ctx context.Context, id int) (*dto.Book, error)
}
