package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otosei-ai/otosei-ai-backend/internal/database/repositories"
)

type CreateUserRequest struct {
	Type          string `json:"type"`
	WalletAddress string `json:"wallet_address"`
}

func HandleUserPost(userRepo *repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		if req.Type != "createUser" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type"})
			return
		}

		user, err := userRepo.CreateIfNotExists(req.WalletAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func HandleUserGet(userRepo *repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Query("walletAddress")
		if address == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing walletAddress query param"})
			return
		}

		user, err := userRepo.GetByWalletAddress(address)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}
