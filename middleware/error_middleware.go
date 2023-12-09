package middleware

import (
	"miniproject/common/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *helper.CustomError:
				c.AbortWithStatusJSON(e.StatusCode, gin.H{
					"error": e.Message,
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			}
		}
	}
}
