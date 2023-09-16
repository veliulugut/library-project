package login

import "github.com/gin-gonic/gin"

type Handler interface {
	Login(c *gin.Context)
}
