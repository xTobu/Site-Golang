package Middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Test use mysql
func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Middleware Test")
		c.Next()
	}
}

//ValidateToken 群组中间件
func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("token")
		if token == "" {
			c.JSON(401, gin.H{
				"message": "Token required",
			})
			// c.Abort()表示请求被终止。
			c.Abort()
			return
		}
		if token != "accesstoken" {
			c.JSON(http.StatusOK, gin.H{
				"message": "Invalid Token",
			})
			// c.Abort()表示请求被终止。
			c.Abort()
			return
		}
		c.Next()
	}
}
