package book

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"library/ent"
	"library/service/book"
	"net/http"
	"strconv"
	"time"
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
// @Param user body book.CreateBookModel true "Create a new book"
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
// @Param name   path string true "name"
// @Success 204 {object} book.GetBookByName
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
// @Param id   path int true "id"
// @Success 200 {object} book.GetBookModel "ok"
// @Router /book/get/{id} [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
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

	var books *book.GetBookModel

	if books, err = b.bookSrv.GetBookByID(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// ListBook godoc
// @Summary get list of books
// @Schemes
// @Description
// @Tags book
// @Accept  json
// @Produce  json
// @Param pageNumber   query int false "pageNumber"
// @Param pageSize  query int false "pageSize"
// @Param isExport query bool false "isExport"
// @Param onlyThisPageExport query bool false "onlyThisPageExport"
// @Param orderBy query  string false "orderBy"
// @Success 200 {object} book.ListBookResponse "ok"
// @Router /books [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (b *Book) ListBook(c *gin.Context) {
	limitQuery := c.DefaultQuery("pageSize", "10")
	offsetQuery := c.DefaultQuery("pageNumber", "0")

	isExport := c.DefaultQuery("isExport", "false")
	onlyThisPageExport := c.DefaultQuery("onlyThisPageExport", "true")

	fetchPartialData, _ := strconv.ParseBool(onlyThisPageExport)
	orderBy := c.DefaultQuery("orderBy", "desc")

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter error"})
		return
	}

	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter error"})
		return
	}
	if isExport == "false" {
		b.ServeJson(c, limit, offset, orderBy)
	} else if isExport == "true" {
		b.ServeExcel(c, limit, offset, orderBy, !fetchPartialData)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid token"})
	}
}

func (b *Book) ServeJson(c *gin.Context, limit, offset int, orderBy string) {
	books, rowCount, err := b.bookSrv.ListBook(c.Request.Context(), limit, offset, orderBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var entBooks []*ent.Book
	for _, getBook := range books {
		entBook := &ent.Book{
			//ID:        getBook.,
			Title:     getBook.Title,
			Author:    getBook.Author,
			Genre:     getBook.Genre,
			Height:    getBook.Height,
			Publisher: getBook.Publisher,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		entBooks = append(entBooks, entBook)
	}

	response := book.ListBookResponse{
		RowCount: rowCount,
		Data:     entBooks,
	}

	c.JSON(http.StatusOK, response)
}

func (b *Book) ServeExcel(c *gin.Context, limit, offset int, orderBy string, fetchAllData bool) {
	f := excelize.NewFile()

	// Set column header
	headers := []string{"ID", "Title", "Author", "Genre", "Height", "Publisher", "CreateAt", "UpdateAt"}
	err := f.SetSheetRow("Sheet1", "A1", headers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	books, err := b.bookSrv.FetchBookData(limit, offset, orderBy, fetchAllData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	for idx, book := range books {
		row := []interface{}{book.ID, book.Title, book.Author, book.Genre, book.Height, book.Publisher, book.CreatedAt, book.UpdatedAt}
		rowIndex := idx + 2 // Start from the second row
		rowStart := fmt.Sprintf("A%d", rowIndex)
		err := f.SetSheetRow("Sheet1", rowStart, row)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	// Save the spreadsheet to a buffer
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Set headers for file download
	downloadName := time.Now().UTC().Format("20060102150405.xlsx")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buf.Bytes())
}
