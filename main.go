package main

import (
	"ewc-backend-go/controllers"
	"ewc-backend-go/database"
	"ewc-backend-go/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// load the config
	LoadAppConfig()
	// connect the database and migrate
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()
	// Initizlize Router
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
			secured.GET("/users", controllers.GetUsers)
			secured.GET("/users/:id", controllers.GetUserById)
			secured.PUT("users/:id", controllers.UpdateUser)
			secured.DELETE("users/:id", controllers.DeleteUser)
		}
	}

	return router
}
