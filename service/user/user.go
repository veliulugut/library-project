package user

import (
	"context"
	"fmt"
	"library/ent"
	"library/pkg/passwd"
	"library/pkg/repository/dto"
	"library/pkg/repository/entadp"
	"library/pkg/validator"
	"time"
)

var _ Service = (*User)(nil)

type User struct {
	pw   passwd.Interface
	repo entadp.RepositoryInterface
}

func New(pw passwd.Interface, repo entadp.RepositoryInterface) *User {
	return &User{
		pw:   pw,
		repo: repo,
	}
}

func (u *User) CreateUser(ctx context.Context, m *CreateUserModel) error {
	var err error

	emailVld := validator.IsEmailValid{Email: m.Email}
	passVld := validator.IsInRange{
		Text: m.Password,
		Max:  32,
		Min:  6,
	}

	if err = validator.ValidateAll(&emailVld, &passVld); err != nil {
		return fmt.Errorf("create user :%w", err)
	}

	var hashedPW string

	if hashedPW, err = u.pw.Generate(m.Password); err != nil {
		return fmt.Errorf("create user / generate password :%w", err)
	}

	userDTO := dto.User{
		Email:    m.Email,
		Password: hashedPW,
		Username: m.UserName,
	}

	if err = u.repo.User().CreateUser(ctx, &userDTO); err != nil {
		return fmt.Errorf("create user / user :%w", err)
	}

	return nil
}

func (u *User) GetUserByID(ctx context.Context, id int) (*GetUserModel, error) {
	var (
		user *dto.User
		err  error
	)

	if user, err = u.repo.User().GetUserByID(ctx, id); err != nil {
		return nil, fmt.Errorf("user / get user by id :%w", err)
	}

	return &GetUserModel{
		Email:     user.Email,
		UserName:  user.Username,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (u *User) DeleteUser(ctx context.Context, id int) error {
	if err := u.repo.User().DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("user service / delete :%w", err)
	}

	return nil
}

func (u *User) UpdateUser(ctx context.Context, id int, c *UpdateUserModel) error {

	userDTO := dto.User{
		Email:     c.Email,
		Password:  c.Password,
		Username:  c.UserName,
		UpdatedAt: c.UpdatedAt,
	}

	if err := u.repo.User().UpdatedUser(ctx, id, &userDTO); err != nil {
		return fmt.Errorf("create user / user :%w", err)
	}

	return nil
}

func (u *User) ListBook(ctx context.Context, limit, offset int, orderBy string) ([]*GetUserModel, int, error) {
	books, count, err := u.repo.User().ListUser(ctx, limit, offset, orderBy)
	if err != nil {
		return nil, 0, fmt.Errorf("book srv / list book: %w", err)
	}

	var userModels []*GetUserModel
	for _, dbUser := range books {
		usersModel := dbUserToGetUserModel(dbUser)
		userModels = append(userModels, usersModel)
	}

	return userModels, count, nil
}

func dbUserToGetUserModel(dbUser *ent.User) *GetUserModel {
	return &GetUserModel{
		Email:     dbUser.Email,
		UserName:  dbUser.Username,
		CreatedAt: time.Now(),
	}
}
