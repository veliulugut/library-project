package entadp

import (
	"context"
	"library/pkg/repository/dto"
)

type RepositoryInterface interface {
	User() UserRepositoryInterface
	Book() BookRepositoryInterface
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, c *dto.User) error
	DeleteUser(ctx context.Context, id int) error
	UpdatedUser(ctx context.Context, id int, c *dto.User) error
	GetUserByID(ctx context.Context, id int) (*dto.User, error)
	GetUserByUserName(ctx context.Context, name string) (*dto.User, error)
	GetUserByEmail(ctx context.Context, email string) (*dto.User, error)

	//ListUser(ctx context.Context, limit, offset int, orderBy string) ([]*dto.User, int, error)
}

type BookRepositoryInterface interface {
	CreateBook(ctx context.Context, c *dto.Book) error
	DeleteBook(ctx context.Context, id int) error
	UpdatedBook(ctx context.Context, id int, c *dto.Book) error
	GetBookByName(ctx context.Context, name string) (*dto.Book, error)
	ListBook(ctx context.Context, limit, offset int, orderBy string) ([]*dto.Book, error)
	GetBookByID(ctx context.Context, id int) (*dto.Book, error)
}
