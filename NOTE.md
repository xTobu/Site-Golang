# Note
> 記錄學習過程的大小事  
___
## 2018-03-09
### JWT
#### 當 axios 掛上 axios.defaults.headers.common\['Authorization'\] ,  
#### GET API時會出現以下錯誤
```
 Request header field header is not allowed by Access-Control-Allow-Headers in preflight response.
```
- 原因：  
https://stackoverflow.com/questions/44245588/how-to-send-authorization-header-with-axios
https://developer.mozilla.org/zh-TW/docs/Web/HTTP/CORS

  
- 解決：  
瀏覽器送出了"preflight" request（預檢請求）（一種 OPTIONS method request），
必須在該API增加OPTIONS responds 並回傳 status code 200，  
且允許 Access-Control-Allow-Headers: Authorization header。   
  
- 最後做法：  
寫CORSMiddleware，並判斷Request.Method，  
若是OPTIONS就回傳204，這樣就不必對會接收到prefight的端口開OPTIONS。
  
- 注意：  
CORSMiddleware必須掛在第一個middleware(router.Use)，  
如果AuthMiddleware在之前，這樣prefight進不去任何一端。

***
## 2018-03-08
### c.Bind、c.BindJSON、json.NewDecoder(c.Request.Body).Decode(&studentData)
#### 當axios使用qs.stringify傳送Content-Type:application/x-www-form-urlencoded時
```
axios.post(
  'http://localhost:8080/auth/login',
  qs.stringify({
    username: payload.username,
    password: payload.password,
  })
)
```
- Data是兩個Form Data (DevTools -> Network)  
```
Form Data(2)  
　username: aa  
　password: bb  
```  
- 在go gin裡接收的方式
```
type StudentReq struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}
var studentData StudentReq
c.Bind(&studentData)
```

#### 另一種axios
```
 axios({
 	method: 'post',
     url: 'https://junxiang.webgene.com.tw/api/Student/Submit',
 	data: {
 		name: 'Fred',
 		email: 'Flintstone',
     },
 	headers: {
    'Content-Type': 'application/x-www-form-urlencoded',
     }
 })
```
- 會是一個Form Data  
```
Form Data  
{username: 'aa', password: 'bb'}
```  

- 在go gin裡接收的方式
```
type StudentReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
var studentData StudentReq
json.NewDecoder(c.Request.Body).Decode(&studentData)
```
