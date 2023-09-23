package reset_password

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	_ "github.com/mattn/go-sqlite3"
	"library/ent"
	"library/ent/enttest"
	"library/service/reset_password"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResetPass_SendResetPasswordValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	w := httptest.NewRecorder()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	ctrl := gomock.NewController(t)

	mockService := reset_password.NewMockServiceResetPass(ctrl)

	mockService.EXPECT().SendResetPasswordEmail(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	s := ResetPass{
		res: mockService,
	}

	router.POST("/auth/reset-password", s.SendResetPasswordValidation)

	w = httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/auth/reset-password", bytes.NewBuffer([]byte(`{"field1": "value1", "field2": "value2"}`)))

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, w.Code)
	}

}
