package Models

//SecretKey go-jwt 金鑰
//Username 帳號
//Password 密碼
const (
	SecretKey = "Have a nice day"
	Username  = "admin"
	Password  = "0000"
)

//UserCredentials 帳號密碼 struct
type UserCredentials struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

//https://github.com/gin-gonic/gin/issues/742
//
//axios qs.stringify x-www-form-urlencoded format, 必須用`form:"username"` , c.Bind(&user),
//
// type Person struct {
// 	Name    string `form:"name" json:"name"`
// 	Address string `form:"address" json:"address"`
// }

// func main() {
// 	route := gin.Default()
// 	route.GET("/testing", startPage)
// 	route.Run(":8085")
// }

// func startPage(c *gin.Context) {
// 	var person Person
// 	if c.Bind(&person) == nil {
// 		log.Println("====== Bind By Query String ======")
// 		log.Println(person.Name)
// 		log.Println(person.Address)
// 	}

// 	if c.BindJSON(&person) == nil {
// 		log.Println("====== Bind By JSON ======")
// 		log.Println(person.Name)
// 		log.Println(person.Address)
// 	}

// 	c.String(200, "Success")
// }
