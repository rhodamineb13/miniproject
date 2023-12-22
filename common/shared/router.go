package shared

import (
	"log"
	"miniproject/common/crypto"
	"miniproject/middleware"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()

	handler, err := NewHandler()
	if err != nil {
		log.Fatalf(err.Error() + ": error in creating handler")
	}

	r.Use(middleware.ErrorHandler())
	r.POST("/sign-up", handler.UserHandler.RegisterUser)
	r.POST("/login", handler.UserHandler.Login)

	user := r.Group("/user", middleware.Authorization(crypto.USER))
	user.POST("/return", handler.BorrowHandler.ReturnBook)
	user.POST("/change-password", handler.UserHandler.ChangePassword)

	books := r.Group("/books")
	books.GET("", handler.BookHandler.GetAllBooks)
	books.GET("/:id", middleware.Authorization(crypto.USER), handler.BookHandler.FindBookByID)
	books.POST("", middleware.Authorization(crypto.ADMIN, crypto.USER), handler.BookHandler.AddNewBook)
	books.POST("/:id/borrow", middleware.Authorization(crypto.USER), handler.BorrowHandler.RequestBorrow)

	return r
}
