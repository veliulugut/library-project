package entadp

import (
	"library/ent"
)

var _ RepositoryInterface = (*Repository)(nil)

func NewRepository(dbClient *ent.Client) *Repository {
	return &Repository{
		user: NewUserRepository(dbClient),
		book: NewBookRepository(dbClient),
	}
}

type Repository struct {
	user UserRepositoryInterface
	book BookRepositoryInterface
}

func (r *Repository) Book() BookRepositoryInterface {
	return r.book
}

func (r *Repository) User() UserRepositoryInterface {
	return r.user
}
