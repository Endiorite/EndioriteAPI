package middleware

import (
	"EndioriteAPI/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func KeyAuth() gin.HandlerFunc {
	expectedAPIKey := config.GetEnv("FULL_ACCESS_API_KEY", "")

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must be in Bearer format"})
			c.Abort()
			return
		}

		apiKey := authHeader[len("Bearer "):]

		if apiKey != expectedAPIKey {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid API Key"})
			c.Abort()
			return
		}

		c.Next()
	}
}
