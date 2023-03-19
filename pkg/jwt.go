package pkg

import (
	"log"
	"time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/GerinAryoPrasetia/go-money-tracking-services/config"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userID, email, name string) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["email"] = email
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Generate encoded token and return it as string
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecretKey))
}
