package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/otosei-ai/otosei-ai-backend/internal/api"
	"github.com/otosei-ai/otosei-ai-backend/internal/api/middleware"
	"github.com/otosei-ai/otosei-ai-backend/internal/cache"
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
	chatRepo := repositories.NewChatRepository(db)

	// Initialize Redis
	redisClient := cache.InitRedis(cfg)

	// Initialize OpenRouter client
	openRouterClient := openrouter.NewClient(cfg.OpenRouterAPIKey, cfg.OpenRouterBaseURL,
		cfg.OpenRouterModel, cfg.OpenRouterFallbacks, cfg.ConfidenceThreshold)

	dependencies := api.Dependencies{
		UserRepo:         userRepo,
		MessageRepo:      messageRepo,
		ChatRepo:         chatRepo,
		OpenRouterClient: openRouterClient,
		RedisClient:      redisClient,
	}

	r := gin.Default()
	protected := r.Group("/api")
	protected.Use(middleware.RequireSession(redisClient, userRepo))

	admin := protected.Group("/admin")
	admin.Use(middleware.RequireAdmin())

	// Register routes
	api.RegisterPublicRoutes(r, dependencies)
	api.RegisterAdminRoutes(admin, dependencies)
	api.RegisterProtectedRoutes(protected, dependencies)

	log.Println("Starting server on port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
