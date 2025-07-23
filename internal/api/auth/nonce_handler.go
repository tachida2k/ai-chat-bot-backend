package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otosei-ai/otosei-ai-backend/internal/services"
)

func NonceHandler(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		wallet := c.Query("walletAddress")
		if wallet == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "walletAddress is required"})
			return
		}

		nonce, err := authService.GenerateAndStoreNonce(wallet)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store nonce"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"nonce": nonce,
		})
	}
}
