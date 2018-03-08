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
	// api := router.Group("/api", Middlewares.Test())
	api := router.Group("/api", Middlewares.ValidateTokenMiddleware())
	{
		api.GET("/get", Controllers.APIStudent)
		api.POST("/post", Controllers.APIInsert)
		api.GET("/checktoken", Controllers.AuthCheckToken)

	}

	//group： auth
	auth := router.Group("/auth", Middlewares.CORSMiddleware())
	{
		auth.POST("/login", Controllers.AuthLogin)

	}

	router.Run(config.Port) // listen and serve on 0.0.0.0:8000

}
