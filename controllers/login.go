package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/middleware"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(c *gin.Context) {
	var input struct {
		Email    *string `json:"email"`
		UserName *string `json:"userName"`
		Password *string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email is provided
	if input.Email != nil && *input.Email != "" {
		collections := []string{"companies", "bcas", "users"}
		var user models.User
		var company models.Company
		var bca models.BCA

		for _, collectionName := range collections {
			collection := db.GetCollection(collectionName)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			filter := bson.M{"email": input.Email}
			var result interface{}
			switch collectionName {
			case "bcas":
				result = &bca
			case "companies":
				result = &company
			case "users":
				result = &user
			}

			err := collection.FindOne(ctx, filter).Decode(result)
			if err == nil {
				// Email exists, call SendLoginOTP
				otpResult, err := middleware.SendLoginOTP(*input.Email)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending OTP"})
					return
				}

				// Update OTP and timestamp in the database
				updateFilter := bson.M{"email": input.Email}
				update := bson.M{
					"$set": bson.M{
						"verificationCode":          otpResult["verificationCode"],
						"verificationCodeTimestamp": time.Now(),
					},
				}
				updateResult := collection.FindOneAndUpdate(ctx, updateFilter, update)
				if err := updateResult.Err(); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating OTP"})
					return
				}

				c.JSON(http.StatusOK, gin.H{"message": "OTP sent to your email", "verificationCode": otpResult["verificationCode"]})
				return
			} else if err.Error() == "mongo: no documents in result" {
				continue
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking email"})
				return
			}
		}

		// Email not found in any collection
		c.JSON(http.StatusNotFound, gin.H{"error": "This email is not registered with us."})
		return
	}

	// Check if UserName and Password are provided
	if input.UserName != nil && *input.UserName != "" && input.Password != nil && *input.Password != "" {
		collections := []string{"companies", "bcas", "users"}
		var user models.User
		var company models.Company
		var bca models.BCA

		for _, collectionName := range collections {
			collection := db.GetCollection(collectionName)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			filter := bson.M{"userName": input.UserName}
			var result interface{}
			switch collectionName {
			case "bcas":
				result = &bca
			case "companies":
				result = &company
			case "users":
				result = &user
			}

			err := collection.FindOne(ctx, filter).Decode(result)
			if err == nil {
				// UserName found, check password
				var hashedPassword string
				var role string
				var id primitive.ObjectID

				switch collectionName {
				case "bcas":
					hashedPassword = *bca.Password
					role = *bca.Role
					id = bca.ID
				case "companies":
					hashedPassword = *company.Password
					role = *company.Role
					id = company.ID
				case "users":
					hashedPassword = *user.Password
					role = *user.Role
					id = user.ID
				}

				err = middleware.ComparePassword(hashedPassword, *input.Password)
				if err == nil {
					// Generate JWT Token
					token, err := middleware.GenerateToken(id.Hex(), role)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
						return
					}

					c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
					return
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
					return
				}
			}
		}

		// UserName not found in any collection
		c.JSON(http.StatusNotFound, gin.H{"error": "This username is not registered with us."})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Email or Username and Password required"})
}
