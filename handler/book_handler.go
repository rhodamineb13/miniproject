package handler

import (
	"miniproject/common/dto"
	"miniproject/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService service.BookService
}

func (b *BookHandler) AddNewBook(c *gin.Context) {
	var book *dto.AddBookDTO

	if err := c.ShouldBindJSON(&book); err != nil {
		_ = c.Error(err)
		return
	}

	if err := b.bookService.AddBook(c, book); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, "new book added")
}

func (b *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := b.bookService.GetAllBooks(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, books)
}

func (b *BookHandler) FindBookByID(c *gin.Context) {
	idString := c.Param("id")
	bookID, err := strconv.Atoi(idString)
	if err != nil {
		_ = c.Error(err)
		return
	}

	book, err := b.bookService.FindBookByID(c, uint(bookID))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}
