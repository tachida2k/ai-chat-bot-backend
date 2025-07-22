package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/otosei-ai/otosei-ai-backend/internal/cache"
	"github.com/otosei-ai/otosei-ai-backend/internal/database/repositories"
)

func RequireAdmin(redisClient *cache.RedisClientWrapper, userRepo *repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid authorization header"})
			return
		}
		sessionID := strings.TrimPrefix(authHeader, "Bearer ")

		walletAddress, err := redisClient.GetSession(sessionID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid session or expired"})
			return
		}

		user, err := userRepo.GetByWalletAddress(walletAddress)
		if err != nil || user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		if user.UserType != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			return
		}

		c.Set("walletAddress", walletAddress)
		c.Set("user", user)

		c.Next()
	}
}
