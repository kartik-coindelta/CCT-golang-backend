package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/middleware"
	"CCT-GOLANG-BACKEND/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateUanCheckStatus updates the status of an uan check by ID
func UpdateUanCheckStatus(c *gin.Context) {
	uanID := c.Param("id")

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

	// Define a struct to capture the incoming update data
	var updateData struct {
		Status string `json:"status"`
	}

	// Bind the JSON body to the struct
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Log the extracted status for debugging
	if updateData.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status cannot be empty"})
		return
	}

	// Prepare the update data to include only the status field
	updateFields := bson.M{"uan.status": updateData.Status}

	// Get the uanCheck collection from the database
	collection := db.GetCollection("uanchecks")

	// Call UpdateItem to perform the update
	result, err := db.UpdateItem(uanID, updateFields, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Uan."})
		return
	}

	// Decode the updated document
	var updatedUanCheck models.UanCheck
	if err := result.Decode(&updatedUanCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated Uan."})
		return
	}

	c.JSON(http.StatusOK, updatedUanCheck)
}
