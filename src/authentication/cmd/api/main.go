// cmd/api/main.go
package main

import (
	"log"
	"os"

	"zippilot/authentication/internal/database"
	"zippilot/authentication/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	if err := database.Connect(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	r := router.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}