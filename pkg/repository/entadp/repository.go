package entadp

import (
	"library/ent"
)

var _ RepositoryInterface = (*Repository)(nil)

func NewRepository(dbClient *ent.Client) *Repository {
	return &Repository{
		user:      NewUserRepository(dbClient),
		book:      NewBookRepository(dbClient),
		resetpass: NewResetPassword(dbClient),
	}
}

type Repository struct {
	user      UserRepositoryInterface
	book      BookRepositoryInterface
	resetpass ResetPasswordInterface
}

func (r *Repository) Book() BookRepositoryInterface {
	return r.book
}

func (r *Repository) User() UserRepositoryInterface {
	return r.user
}

func (r *Repository) ResetPass() ResetPasswordInterface {
	return r.resetpass
}
