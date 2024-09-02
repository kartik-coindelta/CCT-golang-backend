package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateVideoKycCheck handles the HTTP request to update an videoKyc check by ID
func UpdateVideoKycCheck(c *gin.Context) {
	videoKycID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the videoKycCheck collection
	collection := db.GetCollection("videokycchecks")

	// Update the item in the database
	result, err := db.UpdateItem(videoKycID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the VideoKyc."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedVideoKycCheck models.VideoKycCheck
	if err := result.Decode(&updatedVideoKycCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated VideoKyc."})
		return
	}

	c.JSON(http.StatusOK, updatedVideoKycCheck)
}
