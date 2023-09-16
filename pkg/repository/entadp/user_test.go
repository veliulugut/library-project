package entadp

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"library/ent"
	"library/ent/enttest"
	"library/pkg/repository/dto"
	"testing"
	"time"
)

func TestUserRepository_CreateUser(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	userRepo := NewUserRepository(client)

	users := []dto.User{
		{
			Email:     "abcoglu@gmail.com",
			Password:  "abc123",
			Username:  "veli test",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Email:     "abcoglu12@gmail.com",
			Password:  "asdf1234",
			Username:  "sadgfsdfs",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Email:     "test@mail.com",
			Password:  "abc123",
			Username:  "delta",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for i := range users {
		t.Run(fmt.Sprintf("CreateUser_Index:%d", i), func(t *testing.T) {
			if err := userRepo.CreateUser(context.Background(), &users[i]); err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("DuplicateEmail", func(t *testing.T) {
			dpl := dto.User{
				Email:     "abcoglu@gmail.com",
				Password:  "213124314",
				Username:  "json",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := userRepo.CreateUser(context.Background(), &dpl); err != nil {
				t.Error(err)
			}
		})
	})

	t.Run("ContextTimeOut", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*0)
		defer cancel()

		if err := userRepo.CreateUser(ctx, &dto.User{}); err == nil {
			t.Error(err)
		}
	})

	t.Run("ContextCancel", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		cancel()

		if err := userRepo.CreateUser(ctx, &dto.User{}); err == nil {
			t.Error(err)
		}
	})
}

func TestUserRepository_DeleteUser(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	userRepo := NewUserRepository(client)

	user := dto.User{
		Email:     "test@mail.com",
		Password:  "123123",
		Username:  "veli",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := userRepo.CreateUser(context.Background(), &user); err != nil {
		t.Error(err)
	}

	test := []struct {
		description string
		userID      int
		expectedErr error
	}{
		{
			description: "pass",
			userID:      1,
			expectedErr: nil,
		},
		{
			description: "did not passed",
			userID:      31,
			expectedErr: ErrNotFound,
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("DeleteUser_%s", tt.description), func(t *testing.T) {
			err := userRepo.DeleteUser(context.Background(), tt.userID)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error :%v but got error : %v", tt.expectedErr, err)
			}
		})
	}
}

func TestUserRepository_GetUserByID(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	userRepo := NewUserRepository(client)

	user := dto.User{
		Email:     "test@mail.com",
		Password:  "123123",
		Username:  "veli",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := userRepo.CreateUser(context.Background(), &user); err != nil {
		t.Error(err)
	}

	test := []struct {
		description string
		userID      int
		expectErr   error
	}{
		{
			description: "pass",
			userID:      1,
			expectErr:   nil,
		},
		{
			description: "did not passed",
			userID:      31,
			expectErr:   ErrNotFound,
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("GetUserByID_%s", tt.description), func(t *testing.T) {
			_, err := userRepo.GetUserByID(context.Background(), tt.userID)
			if !errors.Is(err, tt.expectErr) {
				t.Errorf("expected error :%v but got error : %v", tt.expectErr, err)
			}
		})

	}
}

func TestUserRepository_UpdatedUser(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	userRepo := NewUserRepository(client)

	_, err := client.User.Create().SetUsername("veli").SetEmail("test@mail.com").SetPassword("test").
		SetCreatedAt(time.Now()).SetUpdatedAt(time.Now()).Save(context.Background())
	if err != nil {
		t.Error(err)
	}

	userUpdate := dto.User{
		Email:     "abc@gmail.com",
		Password:  "3131",
		Username:  "delta",
		UpdatedAt: time.Now(),
	}

	test := []struct {
		title     string
		userID    int
		expectErr error
	}{
		{
			title:     "pass",
			userID:    1,
			expectErr: nil,
		},
		{
			title:     "did not passed",
			userID:    31,
			expectErr: ErrNotFound,
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("GetUserUpdatedUser_%s", tt.title), func(t *testing.T) {
			err := userRepo.UpdatedUser(context.Background(), tt.userID, &userUpdate)
			if !errors.Is(err, tt.expectErr) {
				t.Errorf("expected error :%v but got error : %v", tt.expectErr, err)
			}
		})

	}

}

func TestUserRepository_GetUserByUserName(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	userRepo := NewUserRepository(client)

	user := dto.User{
		Email:     "test@mail.com",
		Password:  "123123",
		Username:  "veli",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := userRepo.CreateUser(context.Background(), &user); err != nil {
		t.Error(err)
	}

	test := []struct {
		description string
		userName    string
		expectErr   error
	}{
		{
			description: "pass",
			userName:    "veli",
			expectErr:   nil,
		},
		{
			description: "did not passed",
			userName:    "abc",
			expectErr:   ErrNotFound,
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("GetUserByUserName_%s", tt.description), func(t *testing.T) {
			_, err := userRepo.GetUserByUserName(context.Background(), tt.userName)
			if !errors.Is(err, tt.expectErr) {
				t.Errorf("expected error :%v but got error : %v", tt.expectErr, err)
			}
		})

	}

}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	userRepo := NewUserRepository(client)

	user := dto.User{
		Email:     "test@mail.com",
		Password:  "123123",
		Username:  "veli",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := userRepo.CreateUser(context.Background(), &user); err != nil {
		t.Error(err)
	}

	test := []struct {
		description string
		email       string
		expectErr   error
	}{
		{
			description: "pass",
			email:       user.Email,
			expectErr:   nil,
		},
		{
			description: "did not passed",
			email:       "abc",
			expectErr:   ErrNotFound,
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("GetUserByEmail_%s", tt.description), func(t *testing.T) {
			_, err := userRepo.GetUserByEmail(context.Background(), tt.email)
			if !errors.Is(err, tt.expectErr) {
				t.Errorf("expected error :%v but got error : %v", tt.expectErr, err)
			}
		})

	}
}
