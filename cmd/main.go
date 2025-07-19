package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/otosei-ai/otosei-ai-backend/internal/api"
	"github.com/otosei-ai/otosei-ai-backend/internal/config"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found")
	}

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	r := gin.Default()

	// Register routes
	api.RegisterRoutes(r)

	log.Println("Starting server on port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
