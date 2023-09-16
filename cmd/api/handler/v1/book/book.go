package book

import (
	"github.com/gin-gonic/gin"
	"library/pkg/repository/dto"
	"library/service/book"
	"net/http"
	"strconv"
)

var _ Handler = (*Book)(nil)

func New(b book.ServiceBook) *Book {
	return &Book{
		bookSrv: b,
	}
}

type Book struct {
	bookSrv book.ServiceBook
}

// CreateBook
// @Summary create book
// @Schemes
// @Description create a book
// @Tags book
// @Accept json
// @Produce json
// @Param user body book.CreateBookModel true "Create a new user"
// @Success 201
// @Router /book [post]
func (b *Book) CreateBook(c *gin.Context) {
	var (
		userDb book.CreateBookModel
		err    error
	)

	if err = c.ShouldBindJSON(&userDb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := b.bookSrv.CreateBook(c.Request.Context(), &userDb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteBook godoc
// @Summary delete book by id
// @Schemes
// @Description
// @Tags book
// @Param id   path int true "id"
// @Success 204
// @Router /book/{id} [delete]
func (b *Book) DeleteBook(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = b.bookSrv.DeleteBook(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNotFound)
}

// UpdatedBook godoc
// @Summary update book by id
// @Schemes
// @Description
// @Tags book
// @Accept  json
// @Produce  json
// @Param user body book.UpdateBookModel  true "Update a  user"
// @Param id   path int true "id"
// @Success 204
// @Router /book/{id} [put]
func (b *Book) UpdatedBook(c *gin.Context) {
	var (
		id  int64
		err error
	)

	idParam := c.Param("id")

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userDb book.UpdateBookModel

	if err = b.bookSrv.UpdatedBook(c.Request.Context(), int(id), &userDb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// GetBookByName godoc
// @Summary get book by name
// @Schemes
// @Description
// @Tags book
// @Accept  json
// @Produce  json
// @Param user body dto.Book  true "get book by name"
// @Param name   path string true "name"
// @Success 204
// @Router /book/{name} [get]
func (b *Book) GetBookByName(c *gin.Context) {
	name := c.Param("name")

	books, err := b.bookSrv.GetBookByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBookByID godoc
// @Summary get book by id
// @Schemes
// @Description
// @Tags book
// @Accept  json
// @Produce  json
// @Param user body dto.Book  true "get book by id"
// @Param name   path int true "id"
// @Success 200
// @Router /book/get/{id} [get]
func (b *Book) GetBookByID(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var books *dto.Book

	if books, err = b.bookSrv.GetBookByID(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
