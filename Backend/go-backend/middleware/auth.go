package middleware

import (
	"net/http"
	"strings"

	"example.com/go-backend/utils"
	"github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		tokenString := c.GetHeader("Authorization")

	
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.Header("WWW-Authenticate", `Bearer realm="Example", error="invalid_token", error_description="Authentication failed or token missing"`)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing or malformed"})
			c.Abort()
			return
		}

	
		tokenString = tokenString[7:]

		
		token, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.Header("WWW-Authenticate", `Bearer realm="Example", error="invalid_token", error_description="The access token expired or is invalid"`)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

	
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["user_id"] == nil {
			c.Header("WWW-Authenticate", `Bearer realm="Example", error="invalid_token", error_description="Invalid token claims"`)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

	
		userID := uint(claims["user_id"].(float64))
		c.Set("user_id", userID)

	
		c.Next()
	}
}
