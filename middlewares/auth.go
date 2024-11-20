package middlewares

import (
	"net/http"
	"strings"
	"todo/service"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		token := context.GetHeader("Authorization")

		if !strings.HasPrefix(token, "Bearer ") {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Authorization header must have Bearer token"})
			return
		}

		authToken := strings.TrimPrefix(token, "Bearer ")

		if authToken == "" {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Authorization token is required"})
			return
		}

		userId, err := service.VerifyToken(authToken)

		if err != nil {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Invalid authorization token!"})
			return
		}

		context.Set("userId", userId)
		context.Next()
	}
}
