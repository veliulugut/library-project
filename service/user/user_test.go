package user

import (
	"context"
	"fmt"
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

func TestUser_DeleteUser(t *testing.T) {
	bc := bcrypt.NewBcrypt("verysecretkey", 10)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		userID         int
		expectedErrMsg string
		mockBehavior   func(mock *entadp.MockUserRepositoryInterface)
	}{
		{
			userID:         1,
			expectedErrMsg: "",
			mockBehavior: func(mock *entadp.MockUserRepositoryInterface) {
				mock.EXPECT().DeleteUser(gomock.Any(), 1).Return(nil)
			},
		},
		{
			userID:         2,
			expectedErrMsg: "User not found",
			mockBehavior: func(mock *entadp.MockUserRepositoryInterface) {
				mock.EXPECT().DeleteUser(gomock.Any(), 2).Return(nil)
			},
		},
	}

	repo := entadp.NewMockRepositoryInterface(ctrl)
	user := New(bc, repo)

	for _, tc := range tests {
		t.Run(fmt.Sprintf("DeleteUser(UserID: %d)", tc.userID), func(t *testing.T) {
			userRepo := entadp.NewMockUserRepositoryInterface(ctrl)
			if tc.mockBehavior != nil {
				tc.mockBehavior(userRepo)
			}
			repo.EXPECT().User().Return(userRepo)

			err := user.DeleteUser(context.Background(), tc.userID)

			if err != nil {
				if tc.expectedErrMsg != "" {
					if err.Error() != tc.expectedErrMsg {
						t.Errorf("Beklenen hata mesajı %q, ancak alınan hata mesajı %q", tc.expectedErrMsg, err.Error())
					}
				} else {
					t.Errorf("Hata beklenmiyordu, ancak hata oluştu: %v", err)
				}
			}
		})
	}
}
