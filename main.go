package main

import (
	"Site-Golang/AppStart"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/appengine"
)

var db *sql.DB

func main() {

	AppStart.RouteConfig()
	appengine.Main()
}
