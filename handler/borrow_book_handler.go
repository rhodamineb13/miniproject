package handler

import (
	"miniproject/common/dto"
	"miniproject/common/helper"
	"miniproject/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BorrowHandler struct {
	borrowService service.BorrowService
}

func (b *BorrowHandler) RequestBorrow(c *gin.Context) {
	user, exists := c.Get("user-id")
	if !exists {
		_ = c.Error(helper.ErrUserUnidentified)
		return
	}

	userID := user.(uint)

	book := c.Param("id")
	bookID, err := strconv.Atoi(book)
	if err != nil {
		_ = c.Error(err)
		return
	}

	req := &dto.BorrowBookRequestDTO{
		UserID: uint(userID),
		BookID: uint(bookID),
	}

	if err := b.borrowService.Borrow(c, req); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data":    req,
		"message": "borrow success",
	})
}

func (b *BorrowHandler) ReturnBook(c *gin.Context) {
	var ret *dto.ReturnBookRequest

	if err := c.ShouldBindJSON(&ret); err != nil {
		_ = c.Error(helper.ErrBookNotFound)
	}

	user, exists := c.Get("user-id")
	if !exists {
		_ = c.Error(helper.ErrUserUnidentified)
		return
	}

	userID := user.(uint)

	ret.UserID = userID

	if err := b.borrowService.ReturnBook(c, ret); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "book successfully returned",
	})

}

func NewBorrowHandler(borrowService service.BorrowService) *BorrowHandler {
	return &BorrowHandler{
		borrowService: borrowService,
	}
}
