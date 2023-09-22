package reset_password

import "github.com/gin-gonic/gin"

type Handler interface {
	SendResetPasswordValidation(c *gin.Context)
}
