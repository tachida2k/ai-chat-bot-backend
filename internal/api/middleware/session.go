package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otosei-ai/otosei-ai-backend/internal/cache"
)

func RequireSession(redisClient *cache.RedisClientWrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.GetHeader("X-Session-ID")
		if sessionID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing session ID"})
			return
		}

		walletAddress, err := redisClient.GetWalletAddress(sessionID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired session"})
			return
		}

		c.Set("walletAddress", walletAddress)
		c.Next()
	}
}
