package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AssignBCA(c *gin.Context) {
	userVerificationRequestID := c.Param("id")
	var request struct {
		BCAInfo models.BCA         `json:"BCAInfo"`
		BCAId   primitive.ObjectID `json:"BCAId"`
	}

	// Parse the request body
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Convert userVerificationRequestID to ObjectID
	objID, err := primitive.ObjectIDFromHex(userVerificationRequestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Connect to the userVerificationRequest collection
	collection := db.GetCollection("userverificationrequest")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the document exists
	var existingDocument models.UserVerificationRequest
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&existingDocument)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch document"})
		}
		return
	}

	// Update BCAInfo and BCAId fields
	update := bson.M{
		"$set": bson.M{
			"BCAInfo": request.BCAInfo,
			"BCAId":   request.BCAId,
		},
	}

	// Apply the update
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "BCAInfo updated successfully"})
}
