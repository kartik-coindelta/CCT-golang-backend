package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/middleware"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetSocialMediaCheckByVerificationID handles the HTTP request to get socialMedia checks by verification ID
func GetSocialMediaCheckByVerificationID(c *gin.Context) {
	verificationID := c.Param("id")

	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing or malformed"})
		return
	}
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	// Validate the token and get claims
	_, err := middleware.ValidateToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}


	// Convert the ID string to a MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(verificationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification ID format"})
		return
	}

	collection := db.GetCollection("socialmediachecks")

	var results []models.SocialMediaCheck
	filter := bson.M{"userVerificationRequestId": objID}

	// Query MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from database"})
		return
	}
	defer cursor.Close(ctx)

	// Loop through cursor and decode each document into the results slice
	for cursor.Next(ctx) {
		var result models.SocialMediaCheck
		if err := cursor.Decode(&result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode database result"})
			return
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor error"})
		return
	}

	c.JSON(http.StatusOK, results)
}
