package book

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	_ "github.com/mattn/go-sqlite3"
	"library/ent"
	"library/ent/enttest"
	"library/service/book"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBook_CreateBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	w := httptest.NewRecorder()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	ctrl := gomock.NewController(t)

	mockService := book.NewMockServiceBook(ctrl)

	mockService.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	s := Book{
		bookSrv: mockService,
	}

	router.POST("/book", s.CreateBook)

	w = httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/book", bytes.NewBuffer([]byte(`{"field1": "value1", "field2": "value2"}`)))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}

}

func TestBook_DeleteBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	w := httptest.NewRecorder()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	ctrl := gomock.NewController(t)

	mockService := book.NewMockServiceBook(ctrl)

	mockService.EXPECT().DeleteBook(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	s := Book{
		bookSrv: mockService,
	}

	router.DELETE("/book/:id", s.DeleteBook)

	w = httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/book/1", bytes.NewBuffer([]byte(`{"field1": "value1", "field2": "value2"}`)))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, w.Code)
	}

}

func TestBook_GetBookByName(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	w := httptest.NewRecorder()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	ctrl := gomock.NewController(t)

	mockService := book.NewMockServiceBook(ctrl)

	mockService.EXPECT().GetBookByName(gomock.Any(), "test").Return(nil, nil).AnyTimes()

	s := Book{
		bookSrv: mockService,
	}

	router.GET("/book/:name", s.GetBookByName)

	w = httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/book/test", bytes.NewBuffer([]byte(`{"field1": "test", "field2": "test2"}`)))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestBook_GetBookByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	w := httptest.NewRecorder()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	ctrl := gomock.NewController(t)

	mockService := book.NewMockServiceBook(ctrl)

	mockService.EXPECT().GetBookByID(gomock.Any(), 2).Return(nil, nil).AnyTimes()

	s := Book{
		bookSrv: mockService,
	}

	router.GET("/book/get/:id", s.GetBookByID)

	w = httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/book/get/2", bytes.NewBuffer([]byte(`{"field1": 2, "field2": "test2"}`)))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}
