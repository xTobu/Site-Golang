# Site
> 基於 Golang 的 Back-End 
>
> 過程中的紀錄：[NOTE](NOTE.md)

**Using**
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

#### Dev
```bash
# serve at localhost:8080
go run ./main.go
```

#### Deploy using Google Cloud SDK
```bash
gcloud app deploy
```
