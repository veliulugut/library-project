package login

import (
	"context"
	"fmt"
	"library/pkg/jwt"
	"library/pkg/passwd"
	"library/pkg/repository/entadp"
)

var _ Service = (*Login)(nil)

func New(r entadp.RepositoryInterface, j jwt.Interface, bc passwd.Interface) *Login {
	return &Login{
		repo: r,
		jwt:  j,
		bc:   bc,
	}
}

type Login struct {
	repo entadp.RepositoryInterface
	jwt  jwt.Interface
	bc   passwd.Interface
}

func (l *Login) Login(ctx context.Context, username, password string) (string, error) {
	user, err := l.repo.User().GetUserByUserName(ctx, username)
	if err != nil {
		return "", fmt.Errorf("login :%w", err)
	}

	err = l.bc.Compare(user.Password, password)
	if err != nil {
		return "", fmt.Errorf("login:%w", err)
	}

	token, err := l.jwt.Generate(int(user.ID))
	if err != nil {
		return "", fmt.Errorf("login :%w", err)
	}

	return token, nil
}
