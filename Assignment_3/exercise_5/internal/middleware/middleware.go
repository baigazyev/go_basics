package middleware

import (
	"bookstore-api/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleRequired checks if the user has the required role
func RoleRequired(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := auth.GetUserFromToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		if user.Role != requiredRole {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
