package api

import (
	"github.com/otosei-ai/otosei-ai-backend/internal/cache"
	"github.com/otosei-ai/otosei-ai-backend/internal/database/repositories"
	"github.com/otosei-ai/otosei-ai-backend/internal/llm/openrouter"
)

type Dependencies struct {
	UserRepo         *repositories.UserRepository
	MessageRepo      *repositories.MessageRepository
	ChatRepo         *repositories.ChatRepository
	OpenRouterClient *openrouter.OpenRouterClient
	RedisClient      *cache.RedisClientWrapper
}
