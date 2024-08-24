package middleware

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func VerifyOTP(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
		OTP   int    `json:"otp"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collections := []string{"companies", "bcas", "users"}
	for _, collectionName := range collections { // Corrected here
		collection := db.GetCollection(collectionName)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		filter := bson.M{"email": input.Email}
		var result interface{}
		var storedOTP *int
		var storedTimestamp *time.Time
		var role string
		var id string

		switch collectionName {
		case "companies":
			var company models.Company
			result = &company
			err := collection.FindOne(ctx, filter).Decode(result)
			if err != nil {
				if err.Error() == "mongo: no documents in result" {
					continue // No document found, continue to the next collection
				}
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving data"})
				return
			}
			storedOTP = company.VerificationCode
			storedTimestamp = company.VerificationCodeTimestamp
			role = *company.Role
			id = company.ID.Hex()

		case "bcas":
			var bca models.BCA
			result = &bca
			err := collection.FindOne(ctx, filter).Decode(result)
			if err != nil {
				if err.Error() == "mongo: no documents in result" {
					continue // No document found, continue to the next collection
				}
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving data"})
				return
			}
			storedOTP = bca.VerificationCode
			storedTimestamp = bca.VerificationCodeTimestamp
			role = *bca.Role
			id = bca.ID.Hex()

			// case "users":
			// 	var user models.User
			// 	result = &user
			// 	err := collection.FindOne(ctx, filter).Decode(result)
			// 	if err != nil {
			// 		if err.Error() == "mongo: no documents in result" {
			// 			continue // No document found, continue to the next collection
			// 		}
			// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving data"})
			// 		return
			// 	}
			// 	storedOTP = user.VerificationCode
			// 	storedTimestamp = user.VerificationCodeTimestamp
			// 	role = *user.Role
			// 	id = user.ID.Hex()
		}

		// Check OTP
		if storedOTP != nil && *storedOTP == input.OTP {
			if storedTimestamp != nil {
				expiryTime := storedTimestamp.Add(5 * time.Minute) // Example expiry of 5 minutes
				if time.Now().Before(expiryTime) {
					// Use the existing GenerateToken function from auth.go
					token, err := GenerateToken(id, role) // Ensure this function is correctly imported
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
						return
					}

					c.JSON(http.StatusOK, gin.H{
						"message": "OTP verified successfully",
						"token":   token,
					})
					return
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "OTP expired"})
					return
				}
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "This email is not registered with us."})
}
