package entadp

import (
	"context"
	"fmt"
	"library/ent"
	"library/ent/user"
	"library/pkg/repository/dto"
	"library/pkg/repository/entadp/helper"
	"time"
)

var _ UserRepositoryInterface = (*UserRepository)(nil)

type UserRepository struct {
	DBClient *ent.Client
}

func NewUserRepository(DBClient *ent.Client) *UserRepository {
	return &UserRepository{
		DBClient: DBClient,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, c *dto.User) error {
	_, err := u.DBClient.User.Create().
		SetPassword(c.Password).
		SetUsername(c.Username).SetEmail(c.Email).
		SetCreatedAt(time.Now()).SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("entadp user / get user by user by email :%w", ErrNotFound)
		}
		return fmt.Errorf("entadp user/ create user :%w", err)
	}

	return nil

}

func (u *UserRepository) DeleteUser(ctx context.Context, id int) error {
	if err := u.DBClient.User.DeleteOneID(id).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("entadp user / get user by user by email :%w", ErrNotFound)
		}
		return fmt.Errorf("entadp user / delete user:%w", err)
	}

	return nil
}

func (u *UserRepository) UpdatedUser(ctx context.Context, id int, c *dto.User) error {
	_, err := u.DBClient.User.UpdateOneID(id).
		SetPassword(c.Password).SetUsername(c.Username).
		SetEmail(c.Email).SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("entadp user / get updated user :%w", ErrNotFound)
		}
		return fmt.Errorf("entadp user / updated user :%w", err)
	}

	return nil
}

func (u *UserRepository) GetUserByID(ctx context.Context, id int) (*dto.User, error) {
	var (
		dbUser *ent.User
		err    error
	)

	if dbUser, err = u.DBClient.User.Get(ctx, id); err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("entadp user / get user by user by email :%w", ErrNotFound)
		}
		return nil, fmt.Errorf("entadp user / get user by id :%w", err)
	}

	return helper.DBUserToDTO(dbUser), nil

}

func (u *UserRepository) GetUserByUserName(ctx context.Context, name string) (*dto.User, error) {
	var (
		dbUser *ent.User
		err    error
	)

	if dbUser, err = u.DBClient.User.Query().Where(user.Username(name)).First(ctx); err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("entadp user / get user by user by email :%w", ErrNotFound)
		}
		return nil, fmt.Errorf("entadp user / get user by user name :%w", err)
	}

	return helper.DBUserToDTO(dbUser), nil
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*dto.User, error) {
	var (
		dbUser *ent.User
		err    error
	)

	if dbUser, err = u.DBClient.User.Query().Where(user.Email(email)).First(ctx); err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("entadp user / get user by user by email :%w", ErrNotFound)
		}
		return nil, fmt.Errorf("entadp user / get user by user by email :%w", err)
	}

	return helper.DBUserToDTO(dbUser), nil
}

/*func (u *UserRepository) ListUser(ctx context.Context, limit, offset int, orderBy string) ([]*dto.User, int, error) {
	var (
		user []*ent.User
		err  error
	)

	count, err := u.DBClient.User.Query().Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("user / count :%w", err)
	}

	if orderBy == "desc" {
		user, err = u.DBClient.User.Query().Order(ent.Desc("created_at")).Offset(offset).Limit(limit).All(ctx)
		if err != nil {
			return nil, 0, fmt.Errorf("list all :%w", err)
		}
	} else {
		user, err = u.DBClient.User.Query().Order(ent.Asc("created_at")).Offset(offset).Limit(limit).All(ctx)
		if err != nil {
			return nil, 0, fmt.Errorf("list all :%w", err)
		}
	}

	return user, count, nil
}*/
