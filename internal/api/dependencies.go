package api

import (
	"github.com/otosei-ai/otosei-ai-backend/internal/services"
)

type Dependencies struct {
	ChatService   *services.ChatService
	AuthService   *services.AuthService
	IntentService *services.IntentService
}
