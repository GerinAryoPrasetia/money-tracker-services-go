package pkg

import (
	"time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userID string) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and return it as string
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your-secret-key"))
}
