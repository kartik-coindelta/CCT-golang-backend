package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateEducationCheck handles the HTTP request to update an education check by ID
func UpdateEducationCheck(c *gin.Context) {
	educationID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the educationCheck collection
	collection := db.GetCollection("educationchecks")

	// Update the item in the database
	result, err := db.UpdateItem(educationID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Education."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedEducationCheck models.EducationCheck
	if err := result.Decode(&updatedEducationCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated Education."})
		return
	}

	c.JSON(http.StatusOK, updatedEducationCheck)
}
