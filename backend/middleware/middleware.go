package middleware

import (
	"ecommerce/tokens"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No authorization provided"})
			c.Abort()
			return
		}

		claims, msg := tokens.ValidateToken(clientToken)
		if msg != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Id)
		c.Next()
	}
}
