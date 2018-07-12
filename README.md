# <img src="https://sendeyo.com/up/d8a757d8eb292ed867e978d3554f0b19.svg" height=50 alt="Site" />



> 基於 Golang 的 Back-End 

> [筆記](NOTE.md)
  
![author](https://img.shields.io/badge/Author-Junxiang-yellow.svg)   
![version](https://img.shields.io/badge/Version-0.0.0-blue.svg)
___

## Development

**Dependencies**
 - Go
 - Gin
 - go-sql-driver/mysql
 - google.golang.org/appengine
 - dgrijalva/jwt-go
 
**Database**
 - MySQL
 - Cloud SQL

**Hosting**  
 - Google App Engine  ([Link](https://xtobu-site.appspot.com))
  
  
  
##  Setup

#### Dev at localhost:8080
```bash
go run ./main.go
```

#### Deploy using Google Cloud SDK
```bash
gcloud app deploy
```  
***  
### golang 的練習  

- 驗證
  - dgrijalva/jwt-go  
- 資料庫
  - go-sql-driver/mysql  
- 發佈
  - google.golang.org/appengine  

**api/login**  

username: `admin`  
password: `0000`  


[version-badge]: https://img.shields.io/badge/version-1.0.0-blue.svg
