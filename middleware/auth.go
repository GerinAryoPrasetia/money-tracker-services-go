package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len("Bearer "):]

		// Parse and validate the JWT token
		claims := &UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				c.AbortWithStatus(http.StatusBadRequest)
			}
			return
		}
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Set the user ID in the request context
		c.Set("id", claims.ID)
		c.Set("name", claims.Name)
		c.Set("email", claims.Email)
		c.Next()
	}
}
