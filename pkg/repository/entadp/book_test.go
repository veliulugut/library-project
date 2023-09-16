package entadp

import (
	"context"
	"errors"
	"fmt"
	"library/ent"
	"library/ent/enttest"
	"library/pkg/repository/dto"
	"testing"
	"time"
)

func TestBookRepository_CreateBook(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	bookRepo := NewBookRepository(client)

	book := []dto.Book{
		{
			ID:        1,
			Author:    "test",
			Genre:     "test",
			Height:    "test",
			Title:     "test",
			Publisher: "test",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Author:    "abc",
			Genre:     "abc",
			Title:     "abc",
			Height:    "abc",
			Publisher: "abc",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for i := range book {
		t.Run(fmt.Sprintf("CreateBook_Index:%d", i), func(t *testing.T) {
			if err := bookRepo.CreateBook(context.Background(), &book[i]); err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("ContextTimeOut", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*0)
		defer cancel()

		if err := bookRepo.CreateBook(ctx, &dto.Book{}); err == nil {
			t.Error(err)
		}
	})

	t.Run("ContextCancel", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		cancel()

		if err := bookRepo.CreateBook(ctx, &dto.Book{}); err == nil {
			t.Error(err)
		}
	})
}

func TestBookRepository_GetBookByName(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	bookRepo := NewBookRepository(client)

	book := []dto.Book{
		{
			ID:        1,
			Author:    "test",
			Genre:     "test",
			Height:    "test",
			Title:     "test",
			Publisher: "test",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	test := []struct {
		description string
		bookId      int
		expectedErr error
	}{
		{
			description: "pass",
			bookId:      1,
			expectedErr: nil,
		},
		{
			description: "did not pass",
			bookId:      2,
			expectedErr: ErrNotFound,
		},
	}

	for i, tt := range test {
		t.Run(fmt.Sprintf("GetBookByName_%d", tt.bookId), func(t *testing.T) {
			_, err := bookRepo.GetBookByName(context.Background(), book[i].Title)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error :%v but got error : %v", tt.expectedErr, err)
			}
		})
	}
}

func TestBookRepository_DeleteBook(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	bookRepo := NewBookRepository(client)

	_, err := client.Book.Create().SetCreatedAt(time.Now()).SetUpdatedAt(time.Now()).SetTitle("test").
		SetHeight("280").SetAuthor("test").SetGenre("test").SetPublisher("test").Save(context.Background())
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		description string
		bookID      int
		expectedErr error
	}{
		{
			description: "pass",
			bookID:      1,
			expectedErr: nil,
		},
		{
			description: "did not pass",
			bookID:      2,
			expectedErr: ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("DeleteBook_%d", tt.bookID), func(t *testing.T) {
			err := bookRepo.DeleteBook(context.Background(), 1)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error :%v but got :%v", tt.expectedErr, err)
			}
		})
	}
}
