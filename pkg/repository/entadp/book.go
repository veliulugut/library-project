package entadp

import (
	"context"
	"fmt"
	"library/ent"
	"library/ent/book"
	"library/pkg/repository/dto"
	"library/pkg/repository/entadp/helper"
	"time"
)

var _ BookRepositoryInterface = (*BookRepository)(nil)

func NewBookRepository(DBClient *ent.Client) *BookRepository {
	return &BookRepository{
		DBClient: DBClient,
	}
}

type BookRepository struct {
	DBClient *ent.Client
}

func (b *BookRepository) CreateBook(ctx context.Context, c *dto.Book) error {
	_, err := b.DBClient.Book.Create().
		SetAuthor(c.Author).SetTitle(c.Title).
		SetGenre(c.Genre).SetHeight(c.Height).
		SetPublisher(c.Publisher).
		SetCreatedAt(time.Now()).SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		if ent.IsNotFound(err); err != nil {
			return fmt.Errorf("entadp user / get user by user by email :%w", ErrNotFound)
		}
		return fmt.Errorf("book / create :%w", err)
	}

	return nil
}

func (b *BookRepository) DeleteBook(ctx context.Context, id int) error {
	if err := b.DBClient.Book.DeleteOneID(id).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("entadp user / get user by user by email :%w", ErrNotFound)
		}
		return fmt.Errorf("book / delete book :%w", err)
	}

	return nil
}

func (b *BookRepository) UpdatedBook(ctx context.Context, id int, c *dto.Book) error {
	_, err := b.DBClient.Book.UpdateOneID(id).
		SetAuthor(c.Author).SetGenre(c.Genre).
		SetHeight(c.Height).SetPublisher(c.Publisher).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("entadp book / updated book :%w", err)
	}

	return nil
}

func (b *BookRepository) GetBookByName(ctx context.Context, name string) (*dto.Book, error) {
	var (
		dbBook *ent.Book
		err    error
	)

	if dbBook, err = b.DBClient.Book.Query().Where(book.Title(name)).First(ctx); err != nil {
		return nil, fmt.Errorf("entadp book / get book by name :%w", err)
	}

	return helper.DBBookToDTO(dbBook), nil
}

func (b *BookRepository) GetBookByID(ctx context.Context, id int) (*dto.Book, error) {
	var (
		dbBook *ent.Book
		err    error
	)

	if dbBook, err = b.DBClient.Book.Get(ctx, id); err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("entadp book / get book by id  :%w", ErrNotFound)
		}
		return nil, fmt.Errorf("get book by id :%w", err)
	}

	return helper.DBBookToDTO(dbBook), nil
}

func (b *BookRepository) ListBook(ctx context.Context, limit, offset int, orderBy string) ([]*ent.Book, int, error) {
	var (
		books []*ent.Book
		err   error
	)

	count, err := b.DBClient.Book.Query().Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("count :%w", err)
	}

	if orderBy == "desc" {
		books, err = b.DBClient.Book.Query().Order(ent.Desc("id")).Limit(limit).Offset(offset).All(ctx)
	} else {
		books, err = b.DBClient.Book.Query().Order(ent.Asc("id")).Limit(limit).Offset(offset).All(ctx)
		if err != nil {
			return nil, 0, fmt.Errorf("list book :%w", err)
		}
	}

	return books, count, nil
}

func (b *BookRepository) FetchBookData(limit, offset int, orderBy string, fetchAll bool) ([]*ent.Book, error) {
	query := b.DBClient.Book.Query()

	if orderBy == "desc" {
		query = query.Order(ent.Desc("id"))
	} else {
		query = query.Order(ent.Asc("id"))
	}

	if !fetchAll {
		query = query.Offset(offset).Limit(limit)
	}

	books, err := query.All(context.Background())
	if err != nil {
		return nil, err
	}

	return books, nil
}
