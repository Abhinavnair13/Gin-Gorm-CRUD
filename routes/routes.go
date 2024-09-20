// routes/routes.go
package routes

import (
	"gin-gorm-crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, bc *controllers.BlogController) {
	router.POST("/createBlog", bc.BlogsCreate)
	router.GET("/getAll", bc.BlogsIndex)
	router.GET("/getById/:id", bc.BlogsGetByID)
	router.PUT("/updateBlog/:id", bc.BlogUpdate)
	router.DELETE("/deleteById/:id", bc.DeleteBlog)
}
