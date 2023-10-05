package user

import (
	"context"
	"github.com/golang/mock/gomock"
	"library/pkg/passwd/bcrypt"
	"library/pkg/repository/entadp"
	"testing"
)

func TestUser_CreateUser(t *testing.T) {
	bc := bcrypt.NewBcrypt("verysecretkey", 10)
	ctrl := gomock.NewController(t)

	userRepo := entadp.NewMockUserRepositoryInterface(ctrl)

	repo := entadp.NewMockRepositoryInterface(ctrl)

	repo.EXPECT().User().Return(userRepo)
	userRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)

	user := New(bc, repo)

	t.Run("CreateUser", func(t *testing.T) {

		err := user.CreateUser(context.Background(), &CreateUserModel{
			Password:        "testsad",
			Email:           "test@mail.com",
			UserName:        "testabc",
			ConfirmPassword: "testsad",
		})

		if err != nil {
			t.Error(err)
		}
	})
}
