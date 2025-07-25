package api

import (
	"github.com/tachida2k/ai-chat-bot-backend/internal/services"
)

type Dependencies struct {
	ChatService   *services.ChatService
	AuthService   *services.AuthService
	IntentService *services.IntentService
}
