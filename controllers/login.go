package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func verifyPasswordWithSecret(hashedPassword, password, secretKey string) error {
	combined := password + secretKey
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(combined))
}

func Login(c *gin.Context) {
	var input models.BCA
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	secretKey := os.Getenv("SECRET_KEY")

	collection := db.GetCollection("BCA")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.BCA

	err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&existingUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	err = verifyPasswordWithSecret(*existingUser.Password, *input.Password, secretKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
