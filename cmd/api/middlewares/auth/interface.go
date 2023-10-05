package auth

import "github.com/gin-gonic/gin"

type Auth interface {
	Check() gin.HandlerFunc
}
