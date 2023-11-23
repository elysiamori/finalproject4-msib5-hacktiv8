package middlewares

import (
	"net/http"

	"github.com/elysiamori/finalproject4-hacktiv8-msib5/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleWares() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, err := token.ExtractTokenID(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Set("user_id", userid)
		c.Next()
	}
}
