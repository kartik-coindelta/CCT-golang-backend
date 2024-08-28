package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetByCompanyId(c *gin.Context) {
	// Get query parameters
	CompanyId := c.Query("CompanyId")
	text := c.Query("text")

	// Ensure BCAId is provided
	if CompanyId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BCAId is required"})
		return
	}

	// Convert BCAId to ObjectID
	companyObjectID, err := primitive.ObjectIDFromHex(CompanyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BCAId"})
		return
	}

	// Build the query with an exact match on BCAId
	query := bson.M{"companyId": companyObjectID}

	// If text is provided, add a name filter with regex
	if text != "" {
		query["name"] = bson.M{"$regex": primitive.Regex{Pattern: text, Options: "i"}}
	}

	fmt.Printf("Constructed query: %+v\n", query)

	// Get MongoDB collection
	collection := db.GetCollection("userverificationrequest")

	// Set up a timeout context for the query
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Retrieve items from the database based on the query
	cursor, err := collection.Find(ctx, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var userverificationrequest []models.UserVerificationRequest
	if err = cursor.All(ctx, &userverificationrequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Retrieved companies: %+v\n", userverificationrequest)

	if len(userverificationrequest) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No companies found"})
		return
	}

	// Return the found companies
	c.JSON(http.StatusOK, userverificationrequest)
}
