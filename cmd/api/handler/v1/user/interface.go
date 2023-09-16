package user

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}
