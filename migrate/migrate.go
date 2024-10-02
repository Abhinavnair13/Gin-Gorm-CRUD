package main

import (
	"gin-gorm-crud/initializers"
	"gin-gorm-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}
func main() {
	initializers.DB.AutoMigrate(&models.Blog{})
	initializers.DB.AutoMigrate(&models.User{})
}
