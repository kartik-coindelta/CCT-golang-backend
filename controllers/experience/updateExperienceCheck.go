package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateExperienceCheck handles the HTTP request to update an experience check by ID
func UpdateExperienceCheck(c *gin.Context) {
	experienceID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the experienceCheck collection
	collection := db.GetCollection("experiencechecks")

	// Update the item in the database
	result, err := db.UpdateItem(experienceID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Experience."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedExperienceCheck models.ExperienceCheck
	if err := result.Decode(&updatedExperienceCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated Experience."})
		return
	}

	c.JSON(http.StatusOK, updatedExperienceCheck)
}
