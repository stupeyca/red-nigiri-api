package database

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service interface {
}

type service struct {
	db *gorm.DB
}

var DB *gorm.DB
var (
	username     = os.Getenv("DB_USERNAME")
	password     = os.Getenv("DB_PASSWORD")
	host         = os.Getenv("DB_HOST")
	port         = os.Getenv("DB_PORT")
	databaseName = os.Getenv("DB_DATABASE")
)

func Connect() Service {
	var err error

	database_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, databaseName)

	DB, err = gorm.Open(postgres.Open(database_url), &gorm.Config{})

	if err != nil {
		fmt.Printf("An error occurred while establishing a connection with the database ❌\n\n")
		panic(err)
	} else {
		fmt.Printf("A connection was successfully established with the database ✅\n\n")
		s := &service{db: DB}
		return s
	}
}
