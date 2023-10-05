package user

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"library/pkg/passwd/bcrypt"
	"library/pkg/repository/entadp"
)

func TestUser_CreateUser(t *testing.T) {
	bc := bcrypt.NewBcrypt("verysecretkey", 10)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tt := []struct {
		name           string
		password       string
		email          string
		userName       string
		confirmPass    string
		expectedErrMsg string
		mockBehavior   func(mock *entadp.MockUserRepositoryInterface)
	}{
		{
			name:           "ValidUser",
			password:       "testsad",
			email:          "test@mail.com",
			userName:       "testabc",
			confirmPass:    "testsad",
			expectedErrMsg: "",
			mockBehavior: func(mock *entadp.MockUserRepositoryInterface) {
				mock.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name:           "DuplicateName",
			password:       "abc123",
			email:          "user@mail.com",
			userName:       "testabc",
			confirmPass:    "abc123",
			expectedErrMsg: "Username already exists",
			mockBehavior: func(mock *entadp.MockUserRepositoryInterface) {
				mock.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}

	repo := entadp.NewMockRepositoryInterface(ctrl)
	user := New(bc, repo)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := entadp.NewMockUserRepositoryInterface(ctrl)
			tc.mockBehavior(userRepo)
			repo.EXPECT().User().Return(userRepo)

			err := user.CreateUser(context.Background(), &CreateUserModel{
				Password:        tc.password,
				Email:           tc.email,
				UserName:        tc.userName,
				ConfirmPassword: tc.confirmPass,
			})

			if err != nil {
				if tc.expectedErrMsg != "" {
					if err.Error() != tc.expectedErrMsg {
						t.Errorf("expected error :%q, but got :%q", tc.expectedErrMsg, err.Error())
					}
				} else {
					t.Errorf("Error not expected, but error occurred: %v", err)
				}
			}
		})
	}
}
