package api

import (
	"github.com/gin-gonic/gin"
	"github.com/otosei-ai/otosei-ai-backend/internal/api/chat"
	"github.com/otosei-ai/otosei-ai-backend/internal/llm/openrouter"
)

func RegisterRoutes(r *gin.Engine, openRouterClient *openrouter.OpenRouterClient) {
	// r.GET("/health", HealthCheckHandler)

	r.POST("/api/chat", chat.ChatHandler(openRouterClient))
	r.POST("/api/chat-stream", chat.ChatStreamHandler(openRouterClient))
}
