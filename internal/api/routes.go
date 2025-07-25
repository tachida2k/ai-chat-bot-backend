package api

import (
	"github.com/gin-gonic/gin"

	"github.com/tachida2k/ai-chat-bot-backend/internal/api/auth"
	"github.com/tachida2k/ai-chat-bot-backend/internal/api/chat"
	"github.com/tachida2k/ai-chat-bot-backend/internal/api/intent"
	"github.com/tachida2k/ai-chat-bot-backend/internal/api/user"
)

func RegisterPublicRoutes(r *gin.Engine, deps Dependencies) {
	r.POST("/api/auth/login", auth.LoginHandler(deps.AuthService))
	r.GET("/api/auth/nonce", auth.NonceHandler(deps.AuthService))
}

func RegisterProtectedRoutes(r *gin.RouterGroup, deps Dependencies) {
	r.POST("/chat", chat.ChatHandler(deps.ChatService))

	r.GET("/get-user", user.GetUserHandler(deps.AuthService))
	r.POST("/intent/classify", intent.ClassifyIntentHandler(deps.IntentService))
}

func RegisterAdminRoutes(r *gin.RouterGroup, deps Dependencies) {
	// TODO:
}
