package shared

import (
	"log"
	"miniproject/middleware"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()

	handler, err := NewHandler()
	if err != nil {
		log.Fatalf("error in creating handler")
	}

	user := r.Group("/user", middleware.ErrorHandler())
	user.POST("/sign-up", handler.UserHandler.RegisterUser)

	return r
}
