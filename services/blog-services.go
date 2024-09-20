package services

import (
	"fmt"
	"gin-gorm-crud/initializers"
	"gin-gorm-crud/models"
)

// BlogService defines the interface for blog operations
type BlogService interface {
	CreateBlog(title, body string) (*models.Blog, error)
	GetAllBlogs() ([]models.Blog, error)
	GetBlogByID(id uint) (*models.Blog, error)
	UpdateBlog(id uint, title, body string) (*models.Blog, error)
	DeleteBlog(id uint) error
}

// blogService is the concrete implementation of BlogService
type blogService struct{}

// NewBlogService creates a new instance of BlogService
func NewBlogService() BlogService {
	return &blogService{}
}

// CreateBlog creates a new blog post
func (s *blogService) CreateBlog(title, body string) (*models.Blog, error) {
	blog := models.Blog{Title: title, Body: body}
	result := initializers.DB.Create(&blog)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create blog: %w", result.Error)
	}
	return &blog, nil
}

// GetAllBlogs retrieves all blog posts
func (s *blogService) GetAllBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	result := initializers.DB.Find(&blogs)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve blogs: %w", result.Error)
	}
	return blogs, nil
}

// GetBlogByID retrieves a blog post by its ID
func (s *blogService) GetBlogByID(id uint) (*models.Blog, error) {
	var blog models.Blog
	result := initializers.DB.First(&blog, id)
	if result.Error != nil {
		return nil, fmt.Errorf("blog not found: %w", result.Error)
	}
	return &blog, nil
}

// UpdateBlog updates an existing blog post
func (s *blogService) UpdateBlog(id uint, title, body string) (*models.Blog, error) {
	var blog models.Blog
	result := initializers.DB.First(&blog, id)
	if result.Error != nil {
		return nil, fmt.Errorf("blog not found: %w", result.Error)
	}

	blog.Title = title
	blog.Body = body
	saveResult := initializers.DB.Save(&blog)
	if saveResult.Error != nil {
		return nil, fmt.Errorf("failed to update blog: %w", saveResult.Error)
	}

	return &blog, nil
}

// DeleteBlog deletes a blog post by its ID
func (s *blogService) DeleteBlog(id uint) error {
	result := initializers.DB.Delete(&models.Blog{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete blog: %w", result.Error)
	}
	return nil
}
