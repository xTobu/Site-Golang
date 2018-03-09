package Middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//CORSMiddleware Header().Set
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS", "204 No Content")
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//Headers Authorization preflight RequestMethod:OPTIONS CORSMiddleware
