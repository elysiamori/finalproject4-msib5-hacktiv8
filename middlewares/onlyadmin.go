package middlewares

import (
	"net/http"

	"github.com/elysiamori/finalproject4-hacktiv8-msib5/token"

	"github.com/gin-gonic/gin"
)

func OnlyAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		role, err := token.ExtractUserRole(c)
		if err != nil || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Access denied for non-admin users"})
			c.Abort()
			return
		}

		c.Next()
	}
}
