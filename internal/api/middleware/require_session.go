package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otosei-ai/otosei-ai-backend/internal/cache"
	"github.com/otosei-ai/otosei-ai-backend/internal/database/repositories"
)

func RequireSession(redisClient *cache.RedisClientWrapper, userRepo *repositories.UserRepository) gin.HandlerFunc {
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

		user, err := userRepo.GetByWalletAddress(walletAddress)
		if err != nil || user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		c.Set("walletAddress", walletAddress)
		c.Set("user", user)
		c.Next()
	}
}
