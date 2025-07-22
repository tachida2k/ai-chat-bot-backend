package auth

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/otosei-ai/otosei-ai-backend/internal/cache"
)

func generateNonce() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func NonceHandler(redisClient *cache.RedisClientWrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		wallet := c.Query("walletAddress")
		if wallet == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "walletAddress is required"})
			return
		}

		nonce, err := generateNonce()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate nonce"})
			return
		}

		err = redisClient.SetNonce(wallet, nonce, 5*time.Minute)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store nonce"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"nonce": nonce,
		})
	}
}
