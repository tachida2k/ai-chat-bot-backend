package services

import (
	"github.com/tachida2k/ai-chat-bot-backend/internal/llm/openrouter"
)

type IntentService struct {
	OpenRouterClient *openrouter.OpenRouterClient
}

func NewIntentService(client *openrouter.OpenRouterClient) *IntentService {
	return &IntentService{OpenRouterClient: client}
}

func (s *IntentService) ClassifyIntent(prompt string) (*openrouter.ClassifyIntentResult, error) {
	return s.OpenRouterClient.ClassifyIntent(prompt)
}
