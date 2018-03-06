package Models

//SecretKey go-jwt 金鑰
const (
	SecretKey = "Have a nice day"
)

//UserCredentials 帳號密碼 struct
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
