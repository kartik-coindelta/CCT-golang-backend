package BCA

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBCAByID(c *gin.Context) {
	// Get the BCA ID from the request parameters
	id := c.Param("id")

	// Convert the ID string to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BCA ID"})
		return
	}

	// Create a context for MongoDB operations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the BCA collection
	collection := db.GetCollection("BCA")

	// Find the BCA document with the given ID
	var bca models.BCA
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&bca)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "BCA not found"})
		return
	}

	// Return the found BCA document
	c.JSON(http.StatusOK, bca)
}
