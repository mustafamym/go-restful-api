package main

import (
	"go-restful-api/controllers"
	"go-restful-api/database"
	"go-restful-api/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	//intialize Database
	database.Connection("host=localhost user=postgres password=postgres dbname=godb port=5432 sslmode=disable ")
	database.Migrate()

	//Initialize Router

	router := initRouter()
	router.Run(":8080")

}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
