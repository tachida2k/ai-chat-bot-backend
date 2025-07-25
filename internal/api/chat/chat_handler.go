package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tachida2k/ai-chat-bot-backend/internal/services"
)

type ChatRequest struct {
	Message string `json:"message" binding:"required"`
	ChatID  *uint  `json:"chat_id,omitempty"`
}

func ChatHandler(
	chatService *services.ChatService,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ChatRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		userIDVal, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found in context"})
			return
		}
		userID, ok := userIDVal.(uint)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user type in context"})
			return
		}

		resp, err := chatService.HandleUserChat(userID, req.ChatID, req.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": resp})
	}
}
