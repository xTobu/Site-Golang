package Controllers

import (
	"Site-Golang/Models"
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

	// var user Models.UserCredentials
	// c.Bind(&user)
	// if c.Bind(&user) == nil {
	// 	log.Println("====== Bind By Query String ======")
	// 	log.Println(user.Username)
	// 	log.Println(user.Password)
	// }

	// if c.BindJSON(&user) == nil {
	// 	log.Println("====== Bind By JSON ======")
	// 	log.Println(user.Username)
	// 	log.Println(user.Password)
	// }

	//判斷body Form
	var user Models.UserCredentials
	// err := json.NewDecoder(c.Request.Body).Decode(&user)
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": user.Username,
		})
		return
	}

	//判斷資料正確
	if (strings.ToLower(user.Username) != Models.Username) || (user.Password != Models.Password) {
		c.JSON(http.StatusForbidden, gin.H{
			"result": "Error logging in",
		})
		return

	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	//有效時間: 24 hour
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	//claims["exp"] = time.Now().Add(time.Second * time.Duration(60)).Unix()

	//建立時間
	claims["iat"] = time.Now().Unix()

	claims["UserID"] = 0
	claims["UserWeight"] = 0

	token.Claims = claims

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Error extracting the key",
		})
		fatal(err)
	}

	tokenString, err := token.SignedString([]byte(Models.SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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
