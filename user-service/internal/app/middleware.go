package app

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
	"user-service/internal/utils"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

// AuthMiddleware validates the JWT token and extracts the user ID
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Remove the "Bearer " prefix from the token string if it exists
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Define the claims struct
		claims := &utils.Claims{}

		// Parse the token and validate its signature using the secret key
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		// Check if the token is valid and there's no parsing error
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store the userID in the context, so it can be accessed in the handler
		c.Set("userID", claims.UserID)

		// Proceed with the request
		c.Next()
	}
}
