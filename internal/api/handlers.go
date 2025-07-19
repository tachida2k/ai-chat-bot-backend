package api

import (
	"github.com/gin-gonic/gin"
	"github.com/otosei-ai/otosei-ai-backend/internal/llm"
)

func RegisterRoutes(r *gin.Engine, openRouterClient *llm.OpenRouterClient) {
	// r.GET("/health", HealthCheckHandler)

	r.POST("/chat", ChatHandler(openRouterClient))
}
