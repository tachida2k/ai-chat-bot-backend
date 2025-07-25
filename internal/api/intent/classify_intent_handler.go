package intent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tachida2k/ai-chat-bot-backend/internal/services"
)

type classifyRequest struct {
	Prompt string `json:"prompt"`
}

func ClassifyIntentHandler(intentService *services.IntentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req classifyRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		result, err := intentService.ClassifyIntent(req.Prompt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
