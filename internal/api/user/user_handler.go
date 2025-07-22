package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otosei-ai/otosei-ai-backend/internal/database/repositories"
)

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
