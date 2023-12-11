package middleware

import (
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
			return
		}

		for _, role := range claims.Role {
			if !utils.CheckRole(role, roles) {
				_ = c.Error(err)
				return
			}
		}
		c.Next()

	}
}
