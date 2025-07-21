package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/otosei-ai/otosei-ai-backend/internal/api"
	"github.com/otosei-ai/otosei-ai-backend/internal/config"
	"github.com/otosei-ai/otosei-ai-backend/internal/database"
	"github.com/otosei-ai/otosei-ai-backend/internal/database/repositories"
	"github.com/otosei-ai/otosei-ai-backend/internal/llm/openrouter"
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

	// Initialize Database
	db := database.InitDB(cfg)
	userRepo := repositories.NewUserRepository(db)
	messageRepo := repositories.NewMessageRepository(db)

	// Initialize OpenRouter client
	openRouterClient := openrouter.NewClient(cfg.OpenRouterAPIKey, cfg.OpenRouterBaseURL,
		cfg.OpenRouterModel, cfg.OpenRouterFallbacks, cfg.ConfidenceThreshold)

	dependencies := api.Dependencies{
		UserRepo:         userRepo,
		MessageRepo:      messageRepo,
		OpenRouterClient: openRouterClient,
	}

	r := gin.Default()

	// Register routes
	api.RegisterRoutes(r, dependencies)

	log.Println("Starting server on port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
