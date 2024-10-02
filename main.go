// main.go
package main

import (
	"gin-gorm-crud/controllers"
	initializers "gin-gorm-crud/initializers"
	"gin-gorm-crud/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// Initialize services and controllers

	blogController := controllers.NewBlogController()

	// Setup routes
	server := routes.NewServer(router)
	server.AddRoutes(blogController)

	router.Run()
}
