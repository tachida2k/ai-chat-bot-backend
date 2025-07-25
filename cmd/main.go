package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tachida2k/ai-chat-bot-backend/internal/api"
	"github.com/tachida2k/ai-chat-bot-backend/internal/api/middleware"
	"github.com/tachida2k/ai-chat-bot-backend/internal/cache"
	"github.com/tachida2k/ai-chat-bot-backend/internal/config"
	"github.com/tachida2k/ai-chat-bot-backend/internal/database"
	"github.com/tachida2k/ai-chat-bot-backend/internal/database/repositories"
	"github.com/tachida2k/ai-chat-bot-backend/internal/llm/openrouter"
	"github.com/tachida2k/ai-chat-bot-backend/internal/services"
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

	// Initialize services
	chatService := services.NewChatService(chatRepo, messageRepo, openRouterClient)
	authService := services.NewAuthService(userRepo, redisClient)

	dependencies := api.Dependencies{
		ChatService: chatService,
		AuthService: authService,
	}

	r := gin.Default()
	protected := r.Group("/api")
	protected.Use(middleware.RequireSession(authService))

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
