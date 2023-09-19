//go:generate mockgen -destination=service_mocks.go -package=user library/service/user Service
package user

import (
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, m *CreateUserModel) error
	GetUserByID(ctx context.Context, id int) (*GetUserModel, error)
	DeleteUser(ctx context.Context, id int) error
	ListBook(ctx context.Context, limit, offset int, orderBy string) ([]*GetUserModel, int, error)
	UpdateUser(ctx context.Context, id int, c *UpdateUserModel) error
}
