// main.go
package main

import (
	"gin-gorm-crud/controllers"
	initializers "gin-gorm-crud/initializers"
	"gin-gorm-crud/routes"
	"gin-gorm-crud/services"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// Initialize services and controllers
	blogService := services.NewBlogService()
	blogController := controllers.NewBlogController(blogService)

	// Setup routes
	routes.SetupRouter(router, blogController)

	router.Run()
}
