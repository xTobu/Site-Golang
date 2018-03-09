package Middlewares

import (
	"Site-Golang/Models"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// c.Request.Header.Get("Authorization")

//ValidateTokenMiddleware 群组中间件
func ValidateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("ValidateTokenMiddleware")
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(Models.SecretKey), nil
			})

		if err == nil {
			if token.Valid {
				c.Next()
			} else {
				claims := token.Claims.(jwt.MapClaims)
				c.JSON(http.StatusUnauthorized, gin.H{
					"UserID":  claims["UserID"],
					"message": "Token is not valid",
				})
				// c.Abort()表示请求被终止。
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized access to this resource",
			})
			// c.Abort()表示请求被终止。
			c.Abort()
			return
		}
	}
}
