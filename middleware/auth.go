package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateToken generates a JWT token for a user
func GenerateToken(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Minute * 50).Unix(), // Token expires in 24 hours
		"role":   role,
	}
	fmt.Println(userID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

// ValidateToken validates a JWT token and returns the claims
func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Log the claims to see the data in the token
		log.Println("JWT Claims:")
		for key, value := range claims {
			log.Printf("%s: %v\n", key, value)
		}
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
