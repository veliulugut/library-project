package entadp

import (
	"context"
	"library/ent"
	"library/pkg/repository/dto"
)

type RepositoryInterface interface {
	User() UserRepositoryInterface
	Book() BookRepositoryInterface
	ResetPass() ResetPasswordInterface
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, c *dto.User) error
	DeleteUser(ctx context.Context, id int) error
	UpdatedUser(ctx context.Context, id int, c *dto.User) error
	GetUserByID(ctx context.Context, id int) (*dto.User, error)
	GetUserByUserName(ctx context.Context, name string) (*dto.User, error)
	GetUserByEmail(ctx context.Context, email string) (*dto.User, error)
	ListUser(ctx context.Context, limit, offset int, orderBy string) ([]*ent.User, int, error)
	UpdatePassword(ctx context.Context, email, password string) error
}

type BookRepositoryInterface interface {
	CreateBook(ctx context.Context, c *dto.Book) error
	DeleteBook(ctx context.Context, id int) error
	UpdatedBook(ctx context.Context, id int, c *dto.Book) error
	GetBookByName(ctx context.Context, name string) (*dto.Book, error)
	ListBook(ctx context.Context, limit, offset int, orderBy string) ([]*ent.Book, int, error)
	GetBookByID(ctx context.Context, id int) (*dto.Book, error)
	FetchBookData(limit, offset int, orderBy string, fetchAll bool) ([]*ent.Book, error)
}

type ResetPasswordInterface interface {
	Create(ctx context.Context, v *dto.ResetPasswordValidation) error
	Get(ctx context.Context, email string, code string) (*dto.ResetPasswordValidation, error)
	Delete(ctx context.Context, email string) error
}
