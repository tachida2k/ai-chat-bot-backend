package api

import (
	"github.com/otosei-ai/otosei-ai-backend/internal/database/repositories"
	"github.com/otosei-ai/otosei-ai-backend/internal/llm/openrouter"
)

type Dependencies struct {
	UserRepo         *repositories.UserRepository
	MessageRepo      *repositories.MessageRepository
	OpenRouterClient *openrouter.OpenRouterClient
}
