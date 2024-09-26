package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectToDB() {
	// Retrieve environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect to db")
	}
}
