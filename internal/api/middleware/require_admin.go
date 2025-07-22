package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otosei-ai/otosei-ai-backend/internal/database/entities"
)

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRaw, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing user in context"})
			return
		}

		user, ok := userRaw.(*entities.User)
		if !ok || user.UserType != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			return
		}

		c.Next()
	}
}
