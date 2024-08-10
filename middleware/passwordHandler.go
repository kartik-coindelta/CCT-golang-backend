package middleware

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

var bcryptSecretKey = []byte(os.Getenv("BCRYPT_SECRET_KEY"))

// HashPassword hashes the password with bcrypt and secret key concatenation.
func HashPassword(password string) (string, error) {
	concatenatedPassword := password + string(bcryptSecretKey)
	bytes, err := bcrypt.GenerateFromPassword([]byte(concatenatedPassword), 14)
	return string(bytes), err
}

// ComparePassword compares the hashed password with the plain text password
func ComparePassword(storedHash, password string) error {
	concatenatedPassword := password + string(bcryptSecretKey)
	return bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(concatenatedPassword))
}
