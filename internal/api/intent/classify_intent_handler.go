package intent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otosei-ai/otosei-ai-backend/internal/llm/openrouter"
)

type classifyRequest struct {
	Prompt string `json:"prompt"`
}

func ClassifyIntentHandler(client *openrouter.OpenRouterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req classifyRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		result, err := client.ClassifyIntent(req.Prompt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
