package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

// Replace this with your secret key
const SecretKey = "your-secret-key"

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	return token, err
}

func IsAuthorized(r *http.Request) bool {
	tokenString := r.Header.Get("Authorization") // Assuming token is sent in the Authorization header
	token, err := ValidateToken(tokenString)
	if err != nil {
		return false
	}

	return token.Valid
}
