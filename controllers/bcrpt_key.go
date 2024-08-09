package controllers

import (
	"os"
)

var (
	// Global variable for bcrypt secret key
	bcryptSecretKey = []byte(os.Getenv("BCRYPT_SECRET_KEY"))
)
