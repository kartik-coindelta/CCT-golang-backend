package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateCriminalCheck handles the HTTP request to update an criminal check by ID
func UpdateCriminalCheck(c *gin.Context) {
	criminalID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the criminalCheck collection
	collection := db.GetCollection("criminalchecks")

	// Update the item in the database
	result, err := db.UpdateItem(criminalID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Criminal data."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedCriminalCheck models.CriminalCheck
	if err := result.Decode(&updatedCriminalCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated Criminal data."})
		return
	}

	c.JSON(http.StatusOK, updatedCriminalCheck)
}
