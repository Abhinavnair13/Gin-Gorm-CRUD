package services

import "gin-gorm-crud/models"

type UserService interface {
	SignUp(email string, password string) (*models.User, error)
	Login(email string, password string) (*models.User, error)
}

// userService is the concrete implementation of userService
// type userService struct{}
