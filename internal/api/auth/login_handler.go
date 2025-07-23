package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otosei-ai/otosei-ai-backend/internal/services"
)

type LoginRequest struct {
	WalletAddress string `json:"walletAddress"`
	Message       string `json:"message"`
	Signature     string `json:"signature"`
}

func LoginHandler(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		sessionID, _, err := authService.HandleLogin(req.WalletAddress, req.Message, req.Signature)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"sessionId":     sessionID,
			"walletAddress": req.WalletAddress,
		})
	}
}
