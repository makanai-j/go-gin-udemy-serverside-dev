package middlewares

import (
	"gin-fleamarket/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService services.IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		user, err := authService.GetUserFromToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
