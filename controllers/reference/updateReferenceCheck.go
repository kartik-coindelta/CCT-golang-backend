package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateReferenceCheck handles the HTTP request to update an reference check by ID
func UpdateReferenceCheck(c *gin.Context) {
	referenceID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the referenceCheck collection
	collection := db.GetCollection("referencechecks")

	// Update the item in the database
	result, err := db.UpdateItem(referenceID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Reference."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedReferenceCheck models.ReferenceCheck
	if err := result.Decode(&updatedReferenceCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated Reference."})
		return
	}

	c.JSON(http.StatusOK, updatedReferenceCheck)
}
