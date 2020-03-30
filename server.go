package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/renishb10/foxg-accesslog-service/controllers"
	"github.com/renishb10/foxg-accesslog-service/middlewares"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	server.LoadHTMLGlob("static/*")

	server.StaticFile("/", "static/home.html")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // Not configured fully

	// Basic routers
	apiRouter := server.Group("/api/v1")
	{
		apiRouter.GET("/:userID1/:userID2", controllers.CheckUsersUnique)
	}

	// Admin routers
	apiAdminRouter := server.Group("/admin")
	{
		apiAdminRouter.POST("/data/seed/:count", controllers.SeedData)
		apiAdminRouter.DELETE("/data/seed", controllers.PurgeData)
	}

	// Handle error response when a route is not defined
	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	server.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
