package user

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	_ "github.com/mattn/go-sqlite3"
	"library/ent"
	"library/ent/enttest"
	"library/service/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUser_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	w := httptest.NewRecorder()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	ctrl := gomock.NewController(t)

	mockService := user.NewMockService(ctrl)

	mockService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	s := User{
		userService: mockService,
	}

	router.POST("/user", s.CreateUser)

	w = httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer([]byte(`{"field1": "test", "field2": "value2"}`)))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}
}

func TestUser_DeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	w := httptest.NewRecorder()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	ctrl := gomock.NewController(t)

	mockService := user.NewMockService(ctrl)

	mockService.EXPECT().DeleteUser(gomock.Any(), 1).Return(nil).AnyTimes()

	s := User{
		userService: mockService,
	}

	router.DELETE("/user/:id", s.DeleteUser)

	w = httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/user/1", bytes.NewBuffer([]byte(`{"field1": 1, "field2": "value2"}`)))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, w.Code)
	}

}

func TestUser_GetUserByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	w := httptest.NewRecorder()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	ctrl := gomock.NewController(t)

	mockService := user.NewMockService(ctrl)

	mockService.EXPECT().GetUserByID(gomock.Any(), 1).Return(nil, nil).AnyTimes()

	s := User{
		userService: mockService,
	}

	router.GET("/user/:id", s.GetUserByID)

	w = httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/user/1", bytes.NewBuffer([]byte(`{"field1": 1, "field2": "value2"}`)))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}
