package auth

import (
	"github.com/gin-gonic/gin"
	"library/pkg/jwt"
	"net/http"
	"strings"
)

var _ Auth = (*Middleware)(nil)

func New(j jwt.Interface) *Middleware {
	return &Middleware{
		jwt: j,
	}
}

type Middleware struct {
	jwt jwt.Interface
}

func (mw *Middleware) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := c.Request.Header.Get("Authorization")
		if a == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "authorization field not found"})
			c.Abort()
			return
		}
		token := strings.Replace(a, "Bearer ", "", 1)
		_, err := mw.jwt.Parse(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
	}
}
