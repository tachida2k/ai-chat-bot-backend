package api

import (
	"github.com/gin-gonic/gin"

	"github.com/otosei-ai/otosei-ai-backend/internal/api/chat"
	user "github.com/otosei-ai/otosei-ai-backend/internal/api/user"
)

func RegisterRoutes(r *gin.Engine, deps Dependencies) {
	// r.GET("/health", HealthCheckHandler)

	r.POST("/api/chat", chat.ChatHandler(deps.OpenRouterClient))
	r.POST("/api/chat-stream", chat.ChatStreamHandler(deps.OpenRouterClient))

	r.POST("/api/user", user.HandleUserPost(deps.UserRepo))
	r.GET("/api/user", user.HandleUserGet(deps.UserRepo))
}
