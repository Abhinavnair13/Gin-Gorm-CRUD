// routes/routes.go
package routes

import (
	"gin-gorm-crud/controllers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

// NewServer creates a new instance of Server.
func NewServer(router *gin.Engine) *Server {
	return &Server{router: router}
}

// AddRoutes adds blog routes to the server.
func (s *Server) AddRoutes(bc *controllers.BlogController) *Server {
	v1 := s.router.Group("/api/v1")

	v1.POST("/createBlog", controllers.BlogsCreate)
	v1.GET("/getAll", controllers.BlogsIndex)
	v1.GET("/getById/:id", controllers.BlogsGetByID)
	v1.PUT("/updateBlog/:id", controllers.BlogUpdate)
	v1.DELETE("/deleteById/:id", controllers.DeleteBlog)

	return s
}
