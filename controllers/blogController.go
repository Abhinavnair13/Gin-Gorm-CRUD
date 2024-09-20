package controllers

import (
	"net/http"
	"strconv"

	"gin-gorm-crud/services"

	"github.com/gin-gonic/gin"
)

// BlogController handles HTTP requests related to blogs
type BlogController struct {
	Service services.BlogService
}

// NewBlogController creates a new instance of BlogController
func NewBlogController(service services.BlogService) *BlogController {
	return &BlogController{Service: service}
}

// BlogsCreate handles the creation of a new blog post
func (bc *BlogController) BlogsCreate(c *gin.Context) {
	var body struct {
		//making title and body required for the body
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog, err := bc.Service.CreateBlog(body.Title, body.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created blog",
		"blog":    blog,
	})
}

// BlogsIndex handles fetching all blog posts
func (bc *BlogController) BlogsIndex(c *gin.Context) {
	blogs, err := bc.Service.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blogs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"blogs": blogs,
	})
}

// BlogsGetByID handles fetching a single blog post by ID
func (bc *BlogController) BlogsGetByID(c *gin.Context) {
	// Get the ID from the URL parameters
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	blog, err := bc.Service.GetBlogByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Respond with the created blog
	c.JSON(http.StatusOK, gin.H{
		"blog": blog,
	})
}

// BlogUpdate handles updating an existing blog post
func (bc *BlogController) BlogUpdate(c *gin.Context) {
	// Get the ID from the URL parameters
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	// Define the expected input structure
	var body struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}

	// Bind JSON input to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use the service to update the blog
	blog, err := bc.Service.UpdateBlog(uint(id), body.Title, body.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return
	}
	// Respond with the created blog
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated blog",
		"blog":    blog,
	})
}

// DeleteBlog handles deleting a blog post by ID
func (bc *BlogController) DeleteBlog(c *gin.Context) {
	// Get the ID from the URL parameters
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	// Use the service to delete the blog
	if err := bc.Service.DeleteBlog(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted blog",
	})
}
