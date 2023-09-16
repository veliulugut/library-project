package user

import (
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, m *CreateUserModel) error
	GetUserByID(ctx context.Context, id int) (*GetUserModel, error)
	//ListUser(ctx context.Context, limit, offset int, orderBy string) ([]*ent.User, int, error)
	DeleteUser(ctx context.Context, id int) error
	//UpdateUser(ctx context.Context, id int, c *UpdateUserModel) (*dto.User, error)
	UpdateUser(ctx context.Context, id int, c *UpdateUserModel) error
}
