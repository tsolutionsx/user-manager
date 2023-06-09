package middlewares

import (
	"ewc-backend-go/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			return
		}

		if err := auth.ValidateToken(tokenString); err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		context.Next()
	}
}
