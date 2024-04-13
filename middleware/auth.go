package middleware

import (
	"log"
	"miniproject/common/crypto"
	"miniproject/common/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorization(roles ...crypto.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		reqToken := strings.Split(bearerToken, " ")[1]

		claims, err := crypto.ValidateToken(reqToken)
		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}

		for _, role := range claims.Role {
			if !utils.CheckRole(role, roles) {
				_ = c.Error(err)
				c.Abort()
				return
			}
		}
		log.Println("ID:", claims.ID)
		log.Println(bearerToken)
		c.Set("user-id", claims.ID)
		c.Next()

	}
}

func AddTwoNumber(x, y int) int {
	return x + y
}