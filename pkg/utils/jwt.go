package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creates a new JWT for a given user ID.
func GenerateJWT(userID uint, secretKey string, expirationHours int) (string, error) {
	// Create the claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * time.Duration(expirationHours)).Unix(),
		"iat":     time.Now().Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and return it as a string
	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}
