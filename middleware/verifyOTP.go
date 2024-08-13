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

	collections := []string{"company", "BCA", "user"}
	for _, collectionName := range collections {
		collection := db.GetCollection(collectionName)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		filter := bson.M{"email": input.Email}
		var result interface{}
		var storedOTP *int
		var storedTimestamp *time.Time

		switch collectionName {
		case "company":
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

		case "BCA":
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

		case "user":
			var user models.User
			result = &user
			err := collection.FindOne(ctx, filter).Decode(result)
			if err != nil {
				if err.Error() == "mongo: no documents in result" {
					continue // No document found, continue to the next collection
				}
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving data"})
				return
			}
			storedOTP = user.VerificationCode
			storedTimestamp = user.VerificationCodeTimestamp
		}

		// Check OTP
		if storedOTP != nil && *storedOTP == input.OTP {
			if storedTimestamp != nil {
				expiryTime := storedTimestamp.Add(5 * time.Minute) // Example expiry of 5 minutes
				if time.Now().Before(expiryTime) {

					c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
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
