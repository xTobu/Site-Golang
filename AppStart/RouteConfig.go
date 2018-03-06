package AppStart

import (
	"Site-Golang/Controllers"
	"Site-Golang/Middlewares"

	"github.com/gin-gonic/gin"

	//mysql
	_ "github.com/go-sql-driver/mysql"
)

// ========== server

//Config struct
type Config struct {
	Port         string
	StaticFolder string
	IndexFile    string
}

//SetDefault Sever data
func (config *Config) SetDefault() {
	config.Port = ":8080"
	config.StaticFolder = "../dist"
	config.IndexFile = "../index.html"
}

////////////////////

// RouteConfig blablaba
func RouteConfig() {
	// set config
	config := Config{}
	config.SetDefault()

	// Creates a default gin router
	router := gin.Default() // Grouping routes

	//group： api
	api := router.Group("/api", Middlewares.Test())
	{
		api.GET("/get", Controllers.Student)
		api.POST("/post", Controllers.Insert)
	}

	//group： auth
	auth := router.Group("/auth")
	{
		auth.GET("/get", Controllers.Student)

	}

	router.Run(config.Port) // listen and serve on 0.0.0.0:8000

}
