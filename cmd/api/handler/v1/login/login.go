package login

import (
	"github.com/gin-gonic/gin"
	"library/service/login"
	"net/http"
)

var _ Handler = (*Login)(nil)

func New(s login.Service) *Login {
	return &Login{
		login: s,
	}
}

type Login struct {
	login login.Service
}

// Login
// @Summary login
// @Schemes
// @Description login
// @Tags auth
// @Accept json
// @Produce json
// @Param user body ReqLogin true "login"
// @Success 200
// @Router /login [post]
func (l *Login) Login(c *gin.Context) {
	var (
		icBody ReqLogin
		err    error
	)

	if err = c.ShouldBindJSON(&icBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var token string

	token, err = l.login.Login(c.Request.Context(), icBody.Username, icBody.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user name or password invalid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
