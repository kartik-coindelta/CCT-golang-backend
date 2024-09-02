package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateUanCheck handles the HTTP request to update an uan check by ID
func UpdateUanCheck(c *gin.Context) {
	uanID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the uanCheck collection
	collection := db.GetCollection("uanchecks")

	// Update the item in the database
	result, err := db.UpdateItem(uanID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Uan."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedUanCheck models.UanCheck
	if err := result.Decode(&updatedUanCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated Uan."})
		return
	}

	c.JSON(http.StatusOK, updatedUanCheck)
}
