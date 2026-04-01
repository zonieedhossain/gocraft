package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/testuser/myapp-test/internal/infrastructure"
	"github.com/testuser/myapp-test/internal/delivery/http"
	"github.com/testuser/myapp-test/internal/repository"
	"github.com/testuser/myapp-test/internal/usecase"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, relying on environment variables")
	}

	// Initialize Infrastructure (DB)
	db, err := infrastructure.NewDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Initialize Layers (Clean Architecture)
	repo := repository.NewRepository(db)
	logic := usecase.NewUsecase(repo)
	
	// Start Web Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on port %s using gin framework\n", port)
	if err := http.StartServer(logic, port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
