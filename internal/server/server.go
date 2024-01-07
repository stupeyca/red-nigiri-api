package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"red-nigiri-api/internal/database"
	"red-nigiri-api/internal/model"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port    int
	db      database.Service
	migrate database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:    port,
		db:      database.Connect(),
		migrate: runMigrations(),
	}

	// Declare Server config.
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func runMigrations() error {
	err := database.DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Printf("An error occurred during the automatic migration process ❌\n\n")
		panic(err)
	} else {
		fmt.Printf("Automatic migrations ran successfully ✅\n\n")
		return err
	}

}
