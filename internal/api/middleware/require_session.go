package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tachida2k/ai-chat-bot-backend/internal/services"
)

func RequireSession(sessionService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.GetHeader("X-Session-ID")
		if sessionID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing session ID"})
			return
		}

		user, err := sessionService.GetUserFromSession(sessionID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("walletAddress", user.WalletAddress)
		c.Set("userId", user.ID)
		c.Set("userType", user.UserType)
		c.Next()
	}
}
