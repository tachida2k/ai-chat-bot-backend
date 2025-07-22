package api

import (
	"github.com/gin-gonic/gin"

	"github.com/otosei-ai/otosei-ai-backend/internal/api/admin"
	"github.com/otosei-ai/otosei-ai-backend/internal/api/auth"
	"github.com/otosei-ai/otosei-ai-backend/internal/api/chat"
	"github.com/otosei-ai/otosei-ai-backend/internal/api/intent"
	"github.com/otosei-ai/otosei-ai-backend/internal/api/user"
)

func RegisterPublicRoutes(r *gin.Engine, deps Dependencies) {
	r.POST("/api/auth/login", auth.LoginHandler(deps.RedisClient, deps.UserRepo))
	r.GET("/api/auth/nonce", auth.NonceHandler(deps.RedisClient))
}

func RegisterProtectedRoutes(r *gin.RouterGroup, deps Dependencies) {
	r.POST("/chat", chat.ChatHandler(deps.OpenRouterClient))
	r.POST("/chat-stream", chat.ChatStreamHandler(deps.OpenRouterClient))

	r.GET("/get-user", user.HandleUserGet(deps.UserRepo))

	r.POST("/intent/classify", intent.ClassifyIntentHandler(deps.OpenRouterClient))
}

func RegisterAdminRoutes(r *gin.RouterGroup, deps Dependencies) {
	r.POST("/create-user", admin.HandleCreateUserPost(deps.UserRepo))
}
