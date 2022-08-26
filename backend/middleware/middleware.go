package middleware

import (
	"backend/myJwt"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()

		authTokenString, err := c.Cookie("token")
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		refreshTokenString, err := c.Cookie("refresh_token")
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		csrf, err := c.Cookie("csrf")
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		err = myJwt.CheckIfValid(authTokenString, refreshTokenString, csrf)

		if err != nil {
			fmt.Println(err)
			c.String(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		c.Next()
		return
	}
}
