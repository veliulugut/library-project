package login

import "context"

type Service interface {
	Login(ctx context.Context, username, password string) (string, error)
}
