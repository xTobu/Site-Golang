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
	router.Use(Middlewares.CORSMiddleware())

	//group： api
	api := router.Group("/api")
	// api.Use(Middlewares.CORSMiddleware())
	api.Use(Middlewares.ValidateTokenMiddleware())
	{
		api.GET("/get", Controllers.APIStudent)
		//api.OPTIONS("/get", Controllers.APIStudent)

		api.POST("/post", Controllers.APIInsert)

		api.GET("/checktoken", Controllers.AuthCheckToken)
		//api.OPTIONS("/checktoken", Controllers.AuthCheckToken)

	}

	//group： auth
	auth := router.Group("/auth")
	// auth.Use(Middlewares.CORSMiddleware())
	{
		auth.POST("/login", Controllers.AuthLogin)
		//auth.OPTIONS("/login", Controllers.AuthLogin)

	}

	router.Run(config.Port) // listen and serve on 0.0.0.0:8000

}
