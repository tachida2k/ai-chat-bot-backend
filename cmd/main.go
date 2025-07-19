package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/otosei-ai/otosei-ai-backend/internal/api"
	"github.com/otosei-ai/otosei-ai-backend/internal/config"
	"github.com/otosei-ai/otosei-ai-backend/internal/llm"
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

	// Initialize OpenRouter client
	openRouterClient := llm.NewOpenRouterClient(cfg.OpenRouterAPIKey, cfg.OpenRouterBaseURL, cfg.OpenRouterModel, cfg.OpenRouterFallbacks)

	r := gin.Default()

	// Register routes
	api.RegisterRoutes(r, openRouterClient)

	log.Println("Starting server on port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
