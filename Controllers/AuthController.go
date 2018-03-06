package Controllers

import (
	"Site-Golang/Models"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//AuthLogin 登入 使用go-jwt取得Token
func AuthLogin(c *gin.Context) {

	var user Models.UserCredentials

	err := json.NewDecoder(c.Request.Body).Decode(&user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": "Error in request",
		})
		return
	}

	if strings.ToLower(user.Username) != "someone" {
		if user.Password != "p@ssword" {
			c.JSON(http.StatusOK, gin.H{
				"result": "Error logging in",
			})
			return
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	//有效時間: 1 hour
	//claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["exp"] = time.Now().Add(time.Second * time.Duration(60)).Unix()

	//建立時間
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": "Error extracting the key",
		})
		fatal(err)
	}

	tokenString, err := token.SignedString([]byte(Models.SecretKey))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": "Error while signing the token",
		})
		fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

//AuthCheckToken 檢查Token時效
func AuthCheckToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"valid":  true,
	})
}
