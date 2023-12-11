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

	books := r.Group("/books")
	books.GET("", handler.BookHandler.GetAllBooks)
	books.GET("/:id", handler.BookHandler.FindBookByID)
	books.POST("", middleware.Authorization(crypto.ADMIN, crypto.USER), handler.BookHandler.AddNewBook)

	return r
}
