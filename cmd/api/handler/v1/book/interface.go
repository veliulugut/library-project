//go:generate mockgen -destination=service_mocks.go -package=book library/cmd/api/handler/v1/book Handler
package book

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
	UpdatedBook(c *gin.Context)
	GetBookByName(c *gin.Context)
	GetBookByID(c *gin.Context)
	ListBook(c *gin.Context)
	ServeExcel(c *gin.Context, limit, offset int, orderBy string, fetchAllData bool)
}
