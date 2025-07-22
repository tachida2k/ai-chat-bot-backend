package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/otosei-ai/otosei-ai-backend/internal/cache"
	"github.com/otosei-ai/otosei-ai-backend/internal/database/repositories"
	"github.com/otosei-ai/otosei-ai-backend/pkg/utils"
)

type LoginRequest struct {
	WalletAddress string `json:"walletAddress"`
	Message       string `json:"message"`
	Signature     string `json:"signature"`
}

func LoginHandler(redisClient *cache.RedisClientWrapper, userRepo *repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		// Get expected nonce from Redis
		expectedNonce, err := redisClient.GetNonce(req.WalletAddress)
		if err != nil || expectedNonce == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or expired nonce"})
			return
		}

		// Verify signature
		recoveredAddress, err := utils.RecoverAddressFromSignature(req.Message, req.Signature)
		if err != nil || !strings.EqualFold(recoveredAddress, req.WalletAddress) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid signature"})
			return
		}

		_, err = userRepo.CreateIfNotExists(req.WalletAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}

		// Create session ID
		sessionID := uuid.New().String()

		// Store session in Redis
		if err := redisClient.SetSession(sessionID, req.WalletAddress, 12*time.Hour); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store session"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"sessionId":     sessionID,
			"walletAddress": req.WalletAddress,
		})
	}
}
