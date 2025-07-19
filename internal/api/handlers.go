package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// r.GET("/health", HealthCheckHandler)

	r.GET("/chat", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Otosei AI",
		})
	})
}
